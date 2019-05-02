package services

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type LanguageService struct {
	langRepo repositories.LanguageRepository
}

func NewLanguageService(langRepo repositories.LanguageRepository) LanguageService {
	return LanguageService{langRepo}
}

func (ls LanguageService) GetAll(ctx *gin.Context) map[string]string {
	return ls.langRepo.GetAll()
}
