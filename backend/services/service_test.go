package services

import (
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/repositories/mock_repositories"
	"testing"
)

func getRepoMocks(t *testing.T) (*gomock.Controller, *mock_repositories.MockColorRepository, *mock_repositories.MockVoteRepository,
		*mock_repositories.MockUserRepository, *mock_repositories.MockLanguageRepository) {
	ctrl := gomock.NewController(t)
	mockColorRepo := mock_repositories.NewMockColorRepository(ctrl)
	mockVoteRepo := mock_repositories.NewMockVoteRepository(ctrl)
	mockUserRepo := mock_repositories.NewMockUserRepository(ctrl)
	mockLangRepo := mock_repositories.NewMockLanguageRepository(ctrl)
	return ctrl, mockColorRepo, mockVoteRepo, mockUserRepo, mockLangRepo
}
