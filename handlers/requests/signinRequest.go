package requests

import (
	"github.com/golodash/galidator/v2"
	"github.com/mathleite/gin-boilerplate-api/pkg/validator"
)

type SigninRequest struct {
	Email    string `json:"email" binding:"required" required:"$field is required"`
	Password string `json:"password" binding:"required" required:"$field is required"`
}

func (s *SigninRequest) GetValidator() galidator.Validator {
	return validator.G.Validator(SigninRequest{})
}
