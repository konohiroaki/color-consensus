package controllers

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/services"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserController_GetIDIfLoggedIn_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService, mockClient := mockUserService(ctrl), mockClient(ctrl)

	userID := "id"
	mockClient.EXPECT().GetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().GetID(gomock.Any()).Return(userID, nil)
	controller := newUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.GetIDIfLoggedIn, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"userID":"%s"}`, userID), response.Body.String())
}

func TestUserController_GetIDIfLoggedIn_NotLoggedIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService, mockClient := mockUserService(ctrl), mockClient(ctrl)

	serviceError := "message from service"
	mockClient.EXPECT().GetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().GetID(gomock.Any()).Return("", errors.New(serviceError))
	controller := newUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.GetIDIfLoggedIn, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assertErrorMessageEqual(t, serviceError, response.Body)
}

func TestUserController_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService, mockClient := mockUserService(ctrl), mockClient(ctrl)

	userID := "id--------------------------------36"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().TryLogin(userID, gomock.Any()).Return(true)
	controller := newUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, userID))))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUserController_Login_FailBind(t *testing.T) {
	userID := "id"
	controller := newUserController(nil, nil)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"user":"%s"}`, userID))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageContains(t, "ID: required", response.Body)
}

func TestUserController_Login_FailBindIDLength(t *testing.T) {
	userID := "id"
	controller := newUserController(nil, nil)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, userID))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageContains(t, "ID: len 36", response.Body)
}

func TestUserController_Login_FailService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService, mockClient := mockUserService(ctrl), mockClient(ctrl)

	userID := "id--------------------------------36"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().TryLogin(userID, gomock.Any()).Return(false)
	controller := newUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, userID))))

	assert.Equal(t, http.StatusUnauthorized, response.Code)
	assertErrorMessageEqual(t, "userID not found in repository", response.Body)
}

func TestUserController_SignUpAndLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService, mockClient := mockUserService(ctrl), mockClient(ctrl)

	userID, nationality, birth, gender := "id", "foo", 1900, "bar"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().SignUpAndLogin(nationality, birth, gender, gomock.Any()).Return(userID, nil)
	controller := newUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.SignUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","birth":%d,"gender":"%s"}`, nationality, birth, gender))))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"userID":"%s"}`, userID), response.Body.String())
}

func TestUserController_SignUpAndLogin_FailBind(t *testing.T) {
	nationality, gender := "foo", "bar"
	controller := newUserController(nil, nil)

	response := getResponseRecorder("", controller.SignUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","gender":"%s"}`, nationality, gender))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageContains(t, "Birth: required", response.Body)
}

func TestUserController_SignUpAndLogin_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService, mockClient := mockUserService(ctrl), mockClient(ctrl)

	nationality, birth, gender, serviceError := "foo", 1900, "bar", "internal server error"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().SignUpAndLogin(nationality, birth, gender, gomock.Any()).Return("", services.NewInternalServerError(serviceError))
	controller := newUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.SignUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","birth":%d,"gender":"%s"}`, nationality, birth, gender))))

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assertErrorMessageEqual(t, serviceError, response.Body)
}

func TestUserController_SignUpAndLogin_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService, mockClient := mockUserService(ctrl), mockClient(ctrl)

	nationality, birth, gender, serviceError := "foo", 1900, "bar", "gender format is not correct"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().SignUpAndLogin(nationality, birth, gender, gomock.Any()).Return("", services.NewValidationError(serviceError))
	controller := newUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.SignUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","birth":%d,"gender":"%s"}`, nationality, birth, gender))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, serviceError, response.Body)
}
