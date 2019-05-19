package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"sync"
)

type LanguageService interface {
	GetAll() map[string]string
}

type languageService struct {
	langRepo repositories.LanguageRepository
}

var (
	languageServiceInstance LanguageService
	languageServiceOnce     sync.Once
)

func GetLanguageService() LanguageService {
	languageServiceOnce.Do(func() {
		languageServiceInstance = newLanguageService(repositories.GetLanguageRepository())
	})
	return languageServiceInstance
}

func newLanguageService(langRepo repositories.LanguageRepository) LanguageService {
	return languageService{langRepo}
}

func (ls languageService) GetAll() map[string]string {
	return ls.langRepo.GetAll()
}
