package requests

import (
	"github.com/golodash/galidator/v2"
	"github.com/mathleite/gin-boilerplate-api/pkg/validator"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" required:"$field is required"`
	Email    string `json:"email" binding:"required" required:"$field is required"`
	Password string `json:"password" binding:"required" required:"$field is required"`
}

func (r *RegisterRequest) GetValidator() galidator.Validator {
	return validator.G.Validator(RegisterRequest{})
}
