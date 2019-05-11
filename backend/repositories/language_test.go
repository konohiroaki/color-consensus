package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanguageRepository_GetAll(t *testing.T) {
	langRepo := NewLanguageRepository()

	actual := langRepo.GetAll()

	assert.Len(t, actual, 62)
	assert.Equal(t, "English", actual["en"])
}

func TestLanguageRepository_IsCodePresent_True(t *testing.T) {
	langRepo := NewLanguageRepository()

	actual := langRepo.IsCodePresent("ja")

	assert.True(t, actual)
}

func TestLanguageRepository_IsCodePresent_False(t *testing.T) {
	langRepo := NewLanguageRepository()

	actual := langRepo.IsCodePresent("aa")

	assert.False(t, actual)
}
