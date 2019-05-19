package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
	"sync"
)

type GenderController struct {
	genderService services.GenderService
}

var (
	genderControllerInstance GenderController
	genderControllerOnce     sync.Once
)

func GetGenderController() GenderController {
	genderControllerOnce.Do(func() {
		genderControllerInstance = newGenderController(services.GetGenderService())
	})
	return genderControllerInstance
}

func newGenderController(genderService services.GenderService) GenderController {
	return GenderController{genderService}
}

func (gc GenderController) GetAll(ctx *gin.Context) {
	genders := gc.genderService.GetAll()
	ctx.JSON(http.StatusOK, genders)
}
