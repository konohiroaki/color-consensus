package services

import (
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/repositories/mock_repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll_Success(t *testing.T) {
	ctrl, mockColoRepo := getColorMock(t)
	defer ctrl.Finish()

	fields := []string{"lang", "name", "code"}
	mockColoRepo.EXPECT().GetAll(fields).Return([]map[string]interface{}{})
	service := NewColorService(mockColoRepo)

	actual := service.GetAll()

	assert.Equal(t, []map[string]interface{}{}, actual)
}

func TestAdd_Success(t *testing.T) {
	ctrl, mockColoRepo := getColorMock(t)
	defer ctrl.Finish()

	mockColoRepo.EXPECT().Add("Lang", "Name", "#ff00ff", "User")
	service := NewColorService(mockColoRepo)

	service.Add("Lang", "Name", "#FF00ff", func() (s string, e error) { return "User", nil })
}

func TestGetNeighbors_Cases(t *testing.T) {
	service := NewColorService(nil)

	testCases := []struct {
		code     string
		size     int
		expected []string
		err      string
	}{
		{"000000", 0, []string{}, "size should be between 1 and 4096"},
		{"000000", 4097, []string{}, "size should be between 1 and 4096"},
		{"000000", 4, []string{"#000000", "#100000", "#001000", "#000010"}, ""},
	}

	for _, testCase := range testCases {
		actual, err := service.GetNeighbors(testCase.code, testCase.size)
		assert.ElementsMatch(t, testCase.expected, actual)
		if testCase.err != "" {
			assert.Equal(t, testCase.err, err.Error())
		}
	}
}

func TestIsValidCodeFormat_Cases(t *testing.T) {
	service := NewColorService(nil)

	testCases := []struct {
		argument string
		expected bool
	}{
		{"#049aDF", true},
		{"049aDF", false},
	}

	for _, testCase := range testCases {
		actual, _ := service.IsValidCodeFormat(testCase.argument)
		assert.Equal(t, testCase.expected, actual)
	}
}

func getColorMock(t *testing.T) (*gomock.Controller, *mock_repositories.MockColorRepository) {
	ctrl := gomock.NewController(t)
	mockColorRepo := mock_repositories.NewMockColorRepository(ctrl)
	return ctrl, mockColorRepo
}
