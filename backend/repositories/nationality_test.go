package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNationalityRepository_GetAll(t *testing.T) {
	nationRepo := NewNationalityRepository()

	actual := nationRepo.GetAll()

	assert.Len(t, actual, 249)
	assert.Equal(t, "Japan", actual["JP"])
}

func TestNationalityRepository_IsCodePresent_True(t *testing.T) {
	nationRepo := NewNationalityRepository()

	actual := nationRepo.IsCodePresent("JP")

	assert.True(t, actual)
}

func TestNationalityRepository_IsCodePresent_False(t *testing.T) {
	nationRepo := NewNationalityRepository()

	actual := nationRepo.IsCodePresent("JJ")

	assert.False(t, actual)
}
