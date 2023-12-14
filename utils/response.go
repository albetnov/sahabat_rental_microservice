package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Responses struct{}

type ValidationErrors struct {
	Field   string
	Tag     string
	Message string
}

var Response = Responses{}

func (r Responses) Validation(c *gin.Context, errors *validator.ValidationErrors) {
	out := make([]ValidationErrors, len(*errors))
	for i, fe := range *errors {
		out[i] = ValidationErrors{fe.Field(), fe.Tag(), fe.Error()}
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Payload invalid!",
		"errors":  out,
	})
}
