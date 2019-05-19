package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
	"sync"
)

type LanguageController struct {
	langService services.LanguageService
}

var (
	langControllerInstance LanguageController
	langControllerOnce     sync.Once
)

func GetLanguageController() LanguageController {
	langControllerOnce.Do(func() {
		langControllerInstance = newLanguageController(services.GetLanguageService())
	})
	return langControllerInstance
}

func newLanguageController(langService services.LanguageService) LanguageController {
	return LanguageController{langService}
}

func (lc LanguageController) GetAll(ctx *gin.Context) {
	languages := lc.langService.GetAll()
	ctx.JSON(http.StatusOK, languages)
}
