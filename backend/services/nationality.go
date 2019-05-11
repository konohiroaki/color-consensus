package services

import "github.com/konohiroaki/color-consensus/backend/repositories"

type NationalityService interface {
	GetAll() map[string]string
}

type nationalityService struct {
	nationRepo repositories.NationalityRepository
}

func NewNationalityService(nationRepo repositories.NationalityRepository) NationalityService {
	return nationalityService{nationRepo}
}

func (s nationalityService) GetAll() map[string]string {
	return s.nationRepo.GetAll()
}
