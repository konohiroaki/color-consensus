package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorService_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorRepo := mockColorRepo(ctrl)

	fields := []string{"category", "name", "code"}
	mockColorRepo.EXPECT().GetAll(fields).Return([]map[string]interface{}{})
	service := newColorService(mockColorRepo, nil)

	actual := service.GetAll()

	assert.Equal(t, []map[string]interface{}{}, actual)
}

func TestColorService_Add_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorRepo, colorCategoryRepo := mockColorRepo(ctrl), mockColorCategoryRepo(ctrl)

	colorCategoryRepo.EXPECT().IsPresent(gomock.Any()).Return(true)
	mockColorRepo.EXPECT().Add("Category", "Name", "#ff00ff", "User").Return(nil)
	service := newColorService(mockColorRepo, colorCategoryRepo)

	err := service.Add("Category", "Name", "#FF00ff", func() (s string, e error) { return "User", nil })

	assert.Nil(t, err)
}

func TestColorService_Add_SuccessShortColorCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorRepo, colorCategoryRepo := mockColorRepo(ctrl), mockColorCategoryRepo(ctrl)

	colorCategoryRepo.EXPECT().IsPresent(gomock.Any()).Return(true)
	mockColorRepo.EXPECT().Add("Category", "Name", "#aabbcc", "User").Return(nil)
	service := newColorService(mockColorRepo, colorCategoryRepo)

	err := service.Add("Category", "Name", "#ABC", func() (s string, e error) { return "User", nil })

	assert.Nil(t, err)
}

func TestColorService_Add_SuccessNewCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorRepo, colorCategoryRepo := mockColorRepo(ctrl), mockColorCategoryRepo(ctrl)

	colorCategoryRepo.EXPECT().IsPresent(gomock.Any()).Return(false)
	colorCategoryRepo.EXPECT().Add("Category", "User").Return(nil)
	mockColorRepo.EXPECT().Add("Category", "Name", "#ff00ff", "User").Return(nil)
	service := newColorService(mockColorRepo, colorCategoryRepo)

	err := service.Add("Category", "Name", "#FF00ff", func() (s string, e error) { return "User", nil })

	assert.Nil(t, err)
}

func TestColorService_Add_FailCategoryAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorRepo, colorCategoryRepo := mockColorRepo(ctrl), mockColorCategoryRepo(ctrl)

	colorCategoryRepo.EXPECT().IsPresent(gomock.Any()).Return(false)
	colorCategoryRepo.EXPECT().Add("Category", "User").Return(fmt.Errorf("error"))
	service := newColorService(mockColorRepo, colorCategoryRepo)

	err := service.Add("Category", "Name", "#FF00ff", func() (s string, e error) { return "User", nil })

	assert.Error(t, err)
}

func TestColorService_Add_FailAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorRepo, colorCategoryRepo := mockColorRepo(ctrl), mockColorCategoryRepo(ctrl)

	colorCategoryRepo.EXPECT().IsPresent(gomock.Any()).Return(false)
	colorCategoryRepo.EXPECT().Add("Category", "User").Return(nil)
	mockColorRepo.EXPECT().Add("Category", "Name", "#ff00ff", "User").Return(fmt.Errorf("error"))
	service := newColorService(mockColorRepo, colorCategoryRepo)

	err := service.Add("Category", "Name", "#FF00ff", func() (s string, e error) { return "User", nil })

	assert.Error(t, err)
}

func TestColorService_GetNeighbors_Cases(t *testing.T) {
	service := newColorService(nil, nil)

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
