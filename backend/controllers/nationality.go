package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
)

type NationalityController struct {
	nationService services.NationalityService
}

func NewNationalityController(nationService services.NationalityService) NationalityController {
	return NationalityController{nationService}
}

func (c NationalityController) GetAll(ctx *gin.Context) {
	nationalities := c.nationService.GetAll()
	ctx.JSON(http.StatusOK, nationalities)
}
