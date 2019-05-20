package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
	"sync"
)

type ColorCategoryController struct {
	colorCategoryService services.ColorCategoryService
}

var (
	colorCategoryControllerInstance ColorCategoryController
	colorCategoryControllerOnce     sync.Once
)

func GetColorCategoryController(env string) ColorCategoryController {
	colorCategoryControllerOnce.Do(func() {
		colorCategoryControllerInstance = newColorCategoryController(services.GetColorCategoryService(env))
	})
	return colorCategoryControllerInstance
}

func newColorCategoryController(colorCategoryService services.ColorCategoryService) ColorCategoryController {
	return ColorCategoryController{colorCategoryService}
}

func (ccc ColorCategoryController) GetAll(ctx *gin.Context) {
	categories := ccc.colorCategoryService.GetAll()

	ctx.JSON(http.StatusOK, categories)
	return
}
