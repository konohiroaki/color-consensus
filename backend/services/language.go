package services

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type LanguageService struct{}

func NewLanguageService() LanguageService {
	return LanguageService{}
}

func (LanguageService) GetAll(ctx *gin.Context) map[string]string {
	return repositories.Language(ctx).GetAll()
}
