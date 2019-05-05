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

func TestLanguageRepository_Get(t *testing.T) {
	langRepo := NewLanguageRepository()

	actual, _ := langRepo.Get("ja")

	assert.Equal(t, "Japanese", actual)
}

func TestLanguageRepository_Get_KeyNotFound(t *testing.T) {
	langRepo := NewLanguageRepository()

	_, actual := langRepo.Get("aa")

	assert.Equal(t, "language key not found", actual.Error())
}
