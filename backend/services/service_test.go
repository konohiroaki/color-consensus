package services

import (
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/repositories/mock_repositories"
)

func mockColorRepo(ctrl *gomock.Controller) *mock_repositories.MockColorRepository {
	return mock_repositories.NewMockColorRepository(ctrl)
}

func mockUserRepo(ctrl *gomock.Controller) *mock_repositories.MockUserRepository {
	return mock_repositories.NewMockUserRepository(ctrl)
}

func mockNationRepo(ctrl *gomock.Controller) *mock_repositories.MockNationalityRepository {
	return mock_repositories.NewMockNationalityRepository(ctrl)
}

func mockGenderRepo(ctrl *gomock.Controller) *mock_repositories.MockGenderRepository {
	return mock_repositories.NewMockGenderRepository(ctrl)
}
