package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type LanguageService interface {
	GetAll() map[string]string
}

type languageService struct {
	langRepo repositories.LanguageRepository
}

func NewLanguageService(langRepo repositories.LanguageRepository) LanguageService {
	return languageService{langRepo}
}

func (ls languageService) GetAll() map[string]string {
	return ls.langRepo.GetAll()
}
