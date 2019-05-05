package controllers

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserController_GetIDIfLoggedIn_Success(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockClient.EXPECT().GetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().GetID(gomock.Any()).Return(userID, nil)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.GetIDIfLoggedIn, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"userID":"%s"}`, userID), response.Body.String())
}

func TestUserController_GetIDIfLoggedIn_NotLoggedIn(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	serviceError := "message from service"
	mockClient.EXPECT().GetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().GetID(gomock.Any()).Return("", errors.New(serviceError))
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.GetIDIfLoggedIn, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assertErrorMessageEqual(t, serviceError, response.Body)
}

func TestUserController_Login_Success(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().TryLogin(userID, gomock.Any()).Return(true)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, userID))))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUserController_Login_FailBind(t *testing.T) {
	userID := "id"
	controller := NewUserController(nil, nil)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"user":"%s"}`, userID))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "userID should be in the request", response.Body)
}

func TestUserController_Login_FailService(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	userID := "id"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().TryLogin(userID, gomock.Any()).Return(false)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, userID))))

	assert.Equal(t, http.StatusUnauthorized, response.Code)
	assertErrorMessageEqual(t, "userID not found in repository", response.Body)
}

func TestUserController_SignUpAndLogin_Success(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	userID, nationality, birth, gender := "id", "foo", 1000, "bar"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().SignUpAndLogin(nationality, birth, gender, gomock.Any()).Return(userID, true)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.SignUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","birth":%d,"gender":"%s"}`, nationality, birth, gender))))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"userID":"%s"}`, userID), response.Body.String())
}

func TestUserController_SignUpAndLogin_FailBind(t *testing.T) {
	nationality, gender := "foo", "bar"
	controller := NewUserController(nil, nil)

	response := getResponseRecorder("", controller.SignUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","gender":"%s"}`, nationality, gender))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "all nationality, birth, gender should be in the request", response.Body)
}

func TestUserController_SignUpAndLogin_FailService(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	nationality, birth, gender := "foo", 1000, "bar"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().SignUpAndLogin(nationality, birth, gender, gomock.Any()).Return("", false)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.SignUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","birth":%d,"gender":"%s"}`, nationality, birth, gender))))

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assertErrorMessageEqual(t, "internal server error", response.Body)
}
