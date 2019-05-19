package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenderRepository_GetAll(t *testing.T) {
	genderRepo := GetGenderRepository()

	actual := genderRepo.GetAll()

	assert.Len(t, actual, 3)
	assert.Equal(t, "Others", actual[2])
}

func TestGenderRepository_IsPresent_True(t *testing.T) {
	genderRepo := GetGenderRepository()

	actual := genderRepo.IsPresent("Male")

	assert.True(t, actual)
}

func TestGenderRepository_IsPresent_False(t *testing.T) {
	genderRepo := GetGenderRepository()

	actual := genderRepo.IsPresent("Foo")

	assert.False(t, actual)
}
