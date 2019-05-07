package services

import "github.com/konohiroaki/color-consensus/backend/repositories"

type GenderService interface {
	GetAll() []string
}

type genderService struct {
	genderRepo repositories.GenderRepository
}

func NewGenderService(genderRepo repositories.GenderRepository) GenderService {
	return genderService{genderRepo}
}

func (ls genderService) GetAll() []string {
	return ls.genderRepo.GetAll()
}
