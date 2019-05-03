package services

import (
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/repositories/mock_repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll_Lang(t *testing.T) {
	ctrl, mockLangRepo := getLanguageMock(t)
	ctrl.Finish()

	mockLangRepo.EXPECT().GetAll().Return(map[string]string{})
	service := NewLanguageService(mockLangRepo)

	actual := service.GetAll()

	assert.Equal(t, map[string]string{}, actual)
}

func getLanguageMock(t *testing.T) (*gomock.Controller, *mock_repositories.MockLanguageRepository) {
	ctrl := gomock.NewController(t)
	mockLangRepo := mock_repositories.NewMockLanguageRepository(ctrl)
	return ctrl, mockLangRepo
}
