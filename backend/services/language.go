package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type LanguageService struct {
	langRepo repositories.LanguageRepository
}

func NewLanguageService(langRepo repositories.LanguageRepository) LanguageService {
	return LanguageService{langRepo}
}

func (ls LanguageService) GetAll() map[string]string {
	return ls.langRepo.GetAll()
}
