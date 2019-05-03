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

	id := "id"
	mockClient.EXPECT().GetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().GetID(gomock.Any()).Return(id, nil)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.GetIDIfLoggedIn, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"userID":"%s"}`, id), response.Body.String())
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

	id := "id"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().TryLogin(id, gomock.Any()).Return(true)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, id))))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUserController_Login_FailBind(t *testing.T) {
	id := "id"
	controller := NewUserController(nil, nil)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"user":"%s"}`, id))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "userID should be in the request", response.Body)
}

func TestUserController_Login_FailService(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	id := "id"
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().TryLogin(id, gomock.Any()).Return(false)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.Login,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, id))))

	assert.Equal(t, http.StatusUnauthorized, response.Code)
	assertErrorMessageEqual(t, "userID not found in repository", response.Body)
}

func TestUserController_SingUpAndLogin_Success(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	id, nationality, gender, birth := "id", "foo", "bar", 1000
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().SingUpAndLogin(nationality, gender, birth, gomock.Any()).Return(id, true)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.SingUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","gender":"%s","birth":%d}`, nationality, gender, birth))))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`"%s"`, id), response.Body.String())
}

func TestUserController_SingUpAndLogin_FailBind(t *testing.T) {
	nationality, gender := "foo", "bar"
	controller := NewUserController(nil, nil)

	response := getResponseRecorder("", controller.SingUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","gender":"%s"}`, nationality, gender))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "all nationality, gender, birth should be in the request", response.Body)
}

func TestUserController_SingUpAndLogin_FailService(t *testing.T) {
	ctrl, _, _, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	nationality, gender, birth := "foo", "bar", 1000
	mockClient.EXPECT().SetUserIDFunc(gomock.Any())
	mockUserService.EXPECT().SingUpAndLogin(nationality, gender, birth, gomock.Any()).Return("", false)
	controller := NewUserController(mockUserService, mockClient)

	response := getResponseRecorder("", controller.SingUpAndLogin,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"nationality":"%s","gender":"%s","birth":%d}`, nationality, gender, birth))))

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assertErrorMessageEqual(t, "internal server error", response.Body)
}
