package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GenderController struct{}

func NewGenderController() GenderController {
	return GenderController{}
}

func (gc GenderController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []string{"Female", "Male", "Others"})
}
