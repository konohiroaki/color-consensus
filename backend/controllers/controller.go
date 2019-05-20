package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"strings"
)

func errorResponse(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"message": message,
		},
	}
}

func getBindErrorMessage(err error) string {
	var message string
	switch err.(type) {
	case validator.ValidationErrors:
		message = validationErrorMessage(err.(validator.ValidationErrors))
	default:
		message = "invalid request"
		log.Println(err)
	}
	return message
}

func validationErrorMessage(error validator.ValidationErrors) string {
	message := "Validation Error:\n"
	for _, v := range error {
		message += fmt.Sprintf("- %s: %s %s; ", v.Field, v.Tag, v.Param)
	}
	return strings.TrimSuffix(message, "; ")
}
