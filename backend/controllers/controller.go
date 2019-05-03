package controllers

import "github.com/gin-gonic/gin"

func errorResponse(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"message": message,
		},
	}
}
