package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
)

type GenderController struct {
	genderService services.GenderService
}

func NewGenderController(genderService services.GenderService) GenderController {
	return GenderController{genderService}
}

func (gc GenderController) GetAll(ctx *gin.Context) {
	genders := gc.genderService.GetAll()
	ctx.JSON(http.StatusOK, genders)
}
