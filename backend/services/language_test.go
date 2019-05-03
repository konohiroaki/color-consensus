package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanguageService_GetAll_Success(t *testing.T) {
	ctrl, _, _, _, mockLangRepo := getRepoMocks(t)
	ctrl.Finish()

	mockLangRepo.EXPECT().GetAll().Return(map[string]string{})
	service := NewLanguageService(mockLangRepo)

	actual := service.GetAll()

	assert.Equal(t, map[string]string{}, actual)
}
