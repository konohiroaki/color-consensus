package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
	"sync"
)

type NationalityController struct {
	nationService services.NationalityService
}

var (
	nationControllerInstance NationalityController
	nationControllerOnce     sync.Once
)

func GetNationalityController() NationalityController {
	nationControllerOnce.Do(func() {
		nationControllerInstance = newNationalityController(services.GetNationalityService())
	})
	return nationControllerInstance
}

func newNationalityController(nationService services.NationalityService) NationalityController {
	return NationalityController{nationService}
}

func (c NationalityController) GetAll(ctx *gin.Context) {
	nationalities := c.nationService.GetAll()
	ctx.JSON(http.StatusOK, nationalities)
}
