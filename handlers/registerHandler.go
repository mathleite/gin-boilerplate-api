package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathleite/gin-boilerplate-api/database"
	"github.com/mathleite/gin-boilerplate-api/handlers/requests"
	"github.com/mathleite/gin-boilerplate-api/pkg/bcrypt"
)

func Register(ctx *gin.Context) {
	request := requests.RegisterRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": request.GetValidator().DecryptErrors(err)})
		return
	}

	newUser := database.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword(request.Password),
	}

	if err := newUser.Save(); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func hashPassword(rawPass string) string {
	pass, err := bcrypt.HashPassword(rawPass)
	if err != nil {
		message := "Fail to hash password"
		log.Fatal(message, err)
	}
	return pass
}
