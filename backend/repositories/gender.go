package repositories

import (
	"sync"
)

type GenderRepository interface {
	GetAll() []string
	IsPresent(string) bool
}

type genderRepository struct {
	genderList []string
}

var (
	genderRepoInstance GenderRepository
	genderRepoOnce     sync.Once
)

func GetGenderRepository() GenderRepository {
	genderRepoOnce.Do(func() {
		repository := newGenderRepository()
		repository.setUpData()
		genderRepoInstance = repository
	})
	return genderRepoInstance
}

func newGenderRepository() *genderRepository {
	return &genderRepository{}
}

func (r genderRepository) GetAll() []string {
	return r.genderList
}

func (r genderRepository) IsPresent(gender string) bool {
	for _, v := range r.genderList {
		if gender == v {
			return true
		}
	}
	return false
}

func (r *genderRepository) setUpData() {
	r.genderList = []string{"Female", "Male", "Others"}
}
