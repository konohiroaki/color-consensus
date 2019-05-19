package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestUserService_IsLoggedIn_True(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mockUserRepo(ctrl)

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := newUserService(mockUserRepo, nil, nil)

	getUserID := func() (s string, e error) { return userID, nil }
	actual := service.IsLoggedIn(getUserID)

	assert.True(t, actual)
}

func TestUserService_IsLoggedIn_NotRegisteredUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mockUserRepo(ctrl)

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(false)
	service := newUserService(mockUserRepo, nil, nil)

	getUserID := func() (s string, e error) { return userID, nil }
	actual := service.IsLoggedIn(getUserID)

	assert.False(t, actual)
}

func TestUserService_IsLoggedIn_NotLoggedIn(t *testing.T) {
	service := newUserService(nil, nil, nil)

	getUserID := func() (s string, e error) { return "", errors.New("error message") }
	actual := service.IsLoggedIn(getUserID)

	assert.False(t, actual)
}

func TestUserService_GetID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mockUserRepo(ctrl)

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := newUserService(mockUserRepo, nil, nil)

	getUserID := func() (s string, e error) { return userID, nil }
	actual, _ := service.GetID(getUserID)

	assert.Equal(t, userID, actual)
}

func TestUserService_GetID_NotLoggedIn(t *testing.T) {
	service := newUserService(nil, nil, nil)

	repoError := fmt.Errorf("user is not logged in")
	getUserID := func() (s string, e error) { return "", repoError }
	_, actual := service.GetID(getUserID)

	assert.Equal(t, repoError, actual)
}

func TestUserService_GetID_NotRegisteredUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mockUserRepo(ctrl)

	userID, err := "id", fmt.Errorf("user is not logged in")
	mockUserRepo.EXPECT().IsPresent(userID).Return(false)
	service := newUserService(mockUserRepo, nil, nil)

	getUserID := func() (s string, e error) { return userID, nil }
	_, actual := service.GetID(getUserID)

	assert.Equal(t, err, actual)
}

func TestUserService_SignUpAndLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo, mockNationRepo, mockGenderRepo := mockUserRepo(ctrl), mockNationRepo(ctrl), mockGenderRepo(ctrl)

	userID, nationality, birth, gender := "id", "foo", 1000, "bar"
	mockGenderRepo.EXPECT().IsPresent(gender).Return(true)
	mockNationRepo.EXPECT().IsCodePresent(nationality).Return(true)
	mockUserRepo.EXPECT().Add(nationality, birth, gender).Return(userID)
	service := newUserService(mockUserRepo, mockNationRepo, mockGenderRepo)

	setUserID := func(string) error { return nil }
	actual, err := service.SignUpAndLogin(nationality, birth, gender, setUserID)

	assert.Equal(t, userID, actual)
	assert.Equal(t, nil, err)
}

func TestUserService_SignUpAndLogin_GenderValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGenderRepo := mockGenderRepo(ctrl)

	nationality, birth, gender := "foo", 1000, "bar"
	mockGenderRepo.EXPECT().IsPresent(gender).Return(false)
	service := newUserService(nil, nil, mockGenderRepo)

	setUserID := func(string) error { return errors.New("error message") }
	actual, err := service.SignUpAndLogin(nationality, birth, gender, setUserID)

	assert.Equal(t, "", actual)
	assert.Equal(t, reflect.TypeOf(&ValidationError{}), reflect.TypeOf(err))
	assert.Equal(t, "gender format is not correct", err.Error())
}

func TestUserService_SignUpAndLogin_NationalityValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockNationRepo, mockGenderRepo := mockNationRepo(ctrl), mockGenderRepo(ctrl)

	nationality, birth, gender := "foo", 1000, "bar"
	mockGenderRepo.EXPECT().IsPresent(gender).Return(true)
	mockNationRepo.EXPECT().IsCodePresent(nationality).Return(false)
	service := newUserService(nil, mockNationRepo, mockGenderRepo)

	setUserID := func(string) error { return errors.New("error message") }
	actual, err := service.SignUpAndLogin(nationality, birth, gender, setUserID)

	assert.Equal(t, "", actual)
	assert.Equal(t, reflect.TypeOf(&ValidationError{}), reflect.TypeOf(err))
	assert.Equal(t, "nationality format is not correct", err.Error())
}

func TestUserService_SignUpAndLogin_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo, mockNationRepo, mockGenderRepo := mockUserRepo(ctrl), mockNationRepo(ctrl), mockGenderRepo(ctrl)

	userID, nationality, birth, gender := "id", "foo", 1000, "bar"
	mockGenderRepo.EXPECT().IsPresent(gender).Return(true)
	mockNationRepo.EXPECT().IsCodePresent(nationality).Return(true)
	mockUserRepo.EXPECT().Add(nationality, birth, gender).Return(userID)
	mockUserRepo.EXPECT().Remove(userID)
	service := newUserService(mockUserRepo, mockNationRepo, mockGenderRepo)

	setUserID := func(string) error { return errors.New("error message") }
	actual, err := service.SignUpAndLogin(nationality, birth, gender, setUserID)

	assert.Equal(t, "", actual)
	assert.Equal(t, reflect.TypeOf(&InternalServerError{}), reflect.TypeOf(err))
}

func TestUserService_TryLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mockUserRepo(ctrl)

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := newUserService(mockUserRepo, nil, nil)

	setUserID := func(string) error { return nil }
	actual := service.TryLogin(userID, setUserID)

	assert.True(t, actual)
}

func TestUserService_TryLogin_NotRegisteredUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mockUserRepo(ctrl)

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(false)
	service := newUserService(mockUserRepo, nil, nil)

	setUserID := func(string) error { return nil }
	actual := service.TryLogin(userID, setUserID)

	assert.False(t, actual)
}

func TestUserService_TryLogin_CookieError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mockUserRepo(ctrl)

	userID := "id"
	mockUserRepo.EXPECT().IsPresent(userID).Return(true)
	service := newUserService(mockUserRepo, nil, nil)

	setUserID := func(string) error { return errors.New("error message") }
	actual := service.TryLogin(userID, setUserID)

	assert.False(t, actual)
}
