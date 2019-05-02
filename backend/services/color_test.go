package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type colorRepositoryMock struct {
	getArguments []string
	addArguments []string
}

func (mock *colorRepositoryMock) GetAll(fields []string) []map[string]interface{} {
	mock.getArguments = append(mock.getArguments, fields...)

	return []map[string]interface{}{}
}

func (mock *colorRepositoryMock) Add(lang, name, code, user string) {
	mock.addArguments = append(mock.addArguments, lang, name, code, user)
}

func TestGetAll(t *testing.T) {
	repoMock := &colorRepositoryMock{}
	service := NewColorService(repoMock)

	service.GetAll()

	assert.Equal(t, repoMock.getArguments, []string{"lang", "name", "code"})
}

func TestAdd(t *testing.T) {
	repoMock := &colorRepositoryMock{}
	service := NewColorService(repoMock)

	service.Add("Lang", "Name", "#FF00ff", "User")

	assert.Equal(t, repoMock.addArguments, []string{"Lang", "Name", "#ff00ff", "User"})
}

func TestGetNeighbors(t *testing.T) {
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

func TestIsValidCodeFormat(t *testing.T) {
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
