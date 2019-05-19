package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"sync"
)

type GenderService interface {
	GetAll() []string
}

type genderService struct {
	genderRepo repositories.GenderRepository
}

var (
	genderServiceInstance GenderService
	genderServiceOnce     sync.Once
)

func GetGenderService() GenderService {
	genderServiceOnce.Do(func() {
		genderServiceInstance = newGenderService(repositories.GetGenderRepository())
	})
	return genderServiceInstance
}

func newGenderService(genderRepo repositories.GenderRepository) GenderService {
	return genderService{genderRepo}
}

func (ls genderService) GetAll() []string {
	return ls.genderRepo.GetAll()
}
