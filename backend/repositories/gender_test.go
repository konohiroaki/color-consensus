package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenderRepository_GetAll(t *testing.T) {
	genderRepo := NewGenderRepository()

	actual := genderRepo.GetAll()

	assert.Len(t, actual, 3)
	assert.Equal(t, "Others", actual[2])
}
