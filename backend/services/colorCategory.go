package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"sync"
)

type ColorCategoryService interface {
	GetAll() []string
}

type colorCategoryService struct {
	colorCategoryRepo repositories.ColorCategoryRepository
}

var (
	colorCategoryServiceInstance ColorCategoryService
	colorCategoryServiceOnce     sync.Once
)

func GetColorCategoryService(env string) ColorCategoryService {
	colorCategoryServiceOnce.Do(func() {
		colorCategoryServiceInstance = newColorCategoryService(repositories.GetColorCategoryRepository(env))
	})
	return colorCategoryServiceInstance
}

func newColorCategoryService(colorRepo repositories.ColorCategoryRepository) colorCategoryService {
	return colorCategoryService{colorRepo}
}

func (ccs colorCategoryService) GetAll() []string {
	return ccs.colorCategoryRepo.GetAll()
}
