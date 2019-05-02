package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
)

type LanguageController struct {
	langService services.LanguageService
}

func NewLanguageController(langService services.LanguageService) LanguageController {
	return LanguageController{langService}
}

func (lc LanguageController) GetAll(ctx *gin.Context) {
	languages := lc.langService.GetAll()
	ctx.JSON(http.StatusOK, languages)
}
