package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathleite/gin-boilerplate-api/database"
	"github.com/mathleite/gin-boilerplate-api/handlers/requests"
	"github.com/mathleite/gin-boilerplate-api/handlers/responses"
	"github.com/mathleite/gin-boilerplate-api/pkg/bcrypt"
	"github.com/mathleite/gin-boilerplate-api/pkg/jwt"
)

func Signin(ctx *gin.Context) {
	request := requests.SigninRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": request.GetValidator().DecryptErrors(err)})
		return
	}

	model := database.User{}
	user, err := model.FindByEmail(request.Email)
	if err != nil || !bcrypt.CheckPasswordHash(request.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	jwt, _ := jwt.CreateToken(user.ID.String())
	ctx.JSON(http.StatusOK, responses.SigninResponse{AccessToken: jwt})
}
