package controllers

import (
	"github.com/gin-gonic/gin"
	repo "github.com/konohiroaki/color-consensus/backend/repositories"
	"net/http"
)

type LanguageController struct{}

func (LanguageController) GetAll(ctx *gin.Context) {
	langRepo := repo.Language(ctx)
	ctx.JSON(http.StatusOK, langRepo.GetAll())
}
