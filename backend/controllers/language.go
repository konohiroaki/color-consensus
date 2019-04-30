package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
)

type LanguageController struct{}

func NewLanguageController() LanguageController {
	return LanguageController{}
}

func (LanguageController) GetAll(ctx *gin.Context) {
	languages := services.Language(ctx).GetAll(ctx)
	ctx.JSON(http.StatusOK, languages)
}
