package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorService_GetAll_Success(t *testing.T) {
	ctrl, mockColoRepo, _, _, _ := getRepoMocks(t)
	defer ctrl.Finish()

	fields := []string{"lang", "name", "code"}
	mockColoRepo.EXPECT().GetAll(fields).Return([]map[string]interface{}{})
	service := NewColorService(mockColoRepo)

	actual := service.GetAll()

	assert.Equal(t, []map[string]interface{}{}, actual)
}

func TestColorService_Add_Success(t *testing.T) {
	ctrl, mockColoRepo, _, _, _ := getRepoMocks(t)
	defer ctrl.Finish()

	mockColoRepo.EXPECT().Add("Lang", "Name", "#ff00ff", "User")
	service := NewColorService(mockColoRepo)

	service.Add("Lang", "Name", "#FF00ff", func() (s string, e error) { return "User", nil })
}

func TestColorService_GetNeighbors_Cases(t *testing.T) {
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

func TestColorService_IsValidCodeFormat_Cases(t *testing.T) {
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