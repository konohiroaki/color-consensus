package services

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_IsLoggedIn_True(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := NewUserService(mockUserRepo)

	getUserID := func() (s string, e error) { return userID, nil }
	actual := service.IsLoggedIn(getUserID)

	assert.True(t, actual)
}

func TestUserService_IsLoggedIn_NotRegisteredUser(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(false)
	service := NewUserService(mockUserRepo)

	getUserID := func() (s string, e error) { return userID, nil }
	actual := service.IsLoggedIn(getUserID)

	assert.False(t, actual)
}

func TestUserService_IsLoggedIn_NotLoggedIn(t *testing.T) {
	service := NewUserService(nil)

	getUserID := func() (s string, e error) { return "", errors.New("error message") }
	actual := service.IsLoggedIn(getUserID)

	assert.False(t, actual)
}

func TestUserService_GetID_Success(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := NewUserService(mockUserRepo)

	getUserID := func() (s string, e error) { return userID, nil }
	actual, _ := service.GetID(getUserID)

	assert.Equal(t, userID, actual)
}

func TestUserService_GetID_NotLoggedIn(t *testing.T) {
	service := NewUserService(nil)

	repoError := fmt.Errorf("user is not logged in")
	getUserID := func() (s string, e error) { return "", repoError }
	_, actual := service.GetID(getUserID)

	assert.Equal(t, repoError, actual)
}

func TestUserService_GetID_NotRegisteredUser(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID, err := "id", fmt.Errorf("user is not logged in")
	mockUserRepo.EXPECT().IsPresent(userID).Return(false)
	service := NewUserService(mockUserRepo)

	getUserID := func() (s string, e error) { return userID, nil }
	_, actual := service.GetID(getUserID)

	assert.Equal(t, err, actual)
}

func TestUserService_SignUpAndLogin_Success(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID, nationality, gender, birth := "id", "foo", "bar", 1000
	mockUserRepo.EXPECT().Add(nationality, gender, birth).Return(userID)
	service := NewUserService(mockUserRepo)

	setUserID := func(string) error { return nil }
	actual, success := service.SignUpAndLogin(nationality, gender, birth, setUserID)

	assert.Equal(t, userID, actual)
	assert.True(t, success)
}

func TestUserService_SignUpAndLogin_CookieError(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID, nationality, gender, birth := "id", "foo", "bar", 1000
	mockUserRepo.EXPECT().Add(nationality, gender, birth).Return(userID)
	mockUserRepo.EXPECT().Remove(userID)
	service := NewUserService(mockUserRepo)

	setUserID := func(string) error { return errors.New("error message") }
	actual, success := service.SignUpAndLogin(nationality, gender, birth, setUserID)

	assert.Equal(t, "", actual)
	assert.False(t, success)
}

func TestUserService_TryLogin_Success(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := NewUserService(mockUserRepo)

	setUserID := func(string) error { return nil }
	actual := service.TryLogin(userID, setUserID)

	assert.True(t, actual)
}

func TestUserService_TryLogin_NotRegisteredUser(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(false)
	service := NewUserService(mockUserRepo)

	setUserID := func(string) error { return nil }
	actual := service.TryLogin(userID, setUserID)

	assert.False(t, actual)
}

func TestUserService_TryLogin_CookieError(t *testing.T) {
	ctrl, _, _, mockUserRepo, _ := getRepoMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := NewUserService(mockUserRepo)

	setUserID := func(string) error { return errors.New("error message") }
	actual := service.TryLogin(userID, setUserID)

	assert.False(t, actual)
}

