package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type languageRepositoryMock struct {
	getAllCalled bool
}

func (mock *languageRepositoryMock) GetAll() map[string]string {
	mock.getAllCalled = true
	return map[string]string{}
}

func (mock *languageRepositoryMock) Get(key string) (string, error) {
	return "", nil
}

func TestGetAll_Lang(t *testing.T) {
	repoMock := &languageRepositoryMock{}
	service := NewLanguageService(repoMock)

	service.GetAll()

	assert.True(t, repoMock.getAllCalled)
}
