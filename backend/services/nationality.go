package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"sync"
)

type NationalityService interface {
	GetAll() map[string]string
}

type nationalityService struct {
	nationRepo repositories.NationalityRepository
}

var (
	nationalityServiceInstance NationalityService
	nationalityServiceOnce     sync.Once
)

func GetNationalityService() NationalityService {
	nationalityServiceOnce.Do(func() {
		nationalityServiceInstance = newNationalityService(repositories.GetNationalityRepository())
	})
	return nationalityServiceInstance
}

func newNationalityService(nationRepo repositories.NationalityRepository) NationalityService {
	return nationalityService{nationRepo}
}

func (s nationalityService) GetAll() map[string]string {
	return s.nationRepo.GetAll()
}
