package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/client/mock_client"
	"github.com/konohiroaki/color-consensus/backend/services/mock_services"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll_Success(t *testing.T) {
	ctrl, mockColorService, _, _ := getMocks(t)
	defer ctrl.Finish()

	lang, name, code := "en", "red", "#ff0000"
	mockColorService.EXPECT().GetAll().Return(
		[]map[string]interface{}{{"lang": lang, "name": name, "code": code}})
	controller := NewColorController(mockColorService, nil, nil)

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`[{"code":"%s","lang":"%s","name":"%s"}]`, code, lang, name), response.Body.String())
}

func TestAdd_Success(t *testing.T) {
	ctrl, mockColorService, mockUserService, mockClient := getMocks(t)
	defer ctrl.Finish()

	lang, name, code := "en", "red", "#ff0000"
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	mockColorService = colorFormatValid(mockColorService)
	mockColorService, mockClient = runAdd(mockColorService, mockClient, lang, name, code)
	controller := NewColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"lang":"%s","name":"%s","code":"%s"}`, lang, name, code))))

	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestAdd_FailAuthorization(t *testing.T) {
	ctrl, mockColorService, mockUserService, mockClient := getMocks(t)
	defer ctrl.Finish()

	mockUserService, mockClient = authorizationFail(mockUserService, mockClient)
	controller := NewColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(`{"lang":"en","name":"red","code":"#ff0000"}`)))

	assert.Equal(t, http.StatusForbidden, response.Code)
	assert.Equal(t, `{"error":{"message":"user need to be logged in to add a color"}}`, response.Body.String())
}

func TestAdd_FailBind(t *testing.T) {
	ctrl, mockColorService, mockUserService, mockClient := getMocks(t)
	defer ctrl.Finish()

	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	controller := NewColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(`{"lang":"en","code":"#ff0000"}`))) // "name" not sent

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, `{"error":{"message":"all language, name, code are necessary"}}`, response.Body.String())
}

func TestAdd_FailColorFormatValidation(t *testing.T) {
	ctrl, mockColorService, mockUserService, mockClient := getMocks(t)
	defer ctrl.Finish()

	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	mockColorService, msg := colorFormatInvalid(mockColorService)
	controller := NewColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(`{"lang":"en","name":"red","code":"ff0000"}`)))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"error":{"message":"color code should match regex: %s"}}`, msg), response.Body.String())
}

func TestGetNeighbors_Success(t *testing.T) {
	ctrl, mockColorService, _, _ := getMocks(t)
	defer ctrl.Finish()

	code, size := "ff0000", 1
	mockColorService.EXPECT().GetNeighbors(code, size).Return([]string{"#ff0000"}, nil)
	controller := NewColorController(mockColorService, nil, nil)

	response := getResponseRecorder("/:code", controller.GetNeighbors,
		http.MethodPost, fmt.Sprintf("/%s?size=%d", code, size), nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `["#ff0000"]`, response.Body.String())
}

func TestGetNeighbors_FailSizeAtoiConversion(t *testing.T) {
	ctrl, _, _, _ := getMocks(t)
	defer ctrl.Finish()

	code, size := "ff0000", "a"
	controller := NewColorController(nil, nil, nil)

	response := getResponseRecorder("/:code", controller.GetNeighbors,
		http.MethodPost, fmt.Sprintf("/%s?size=%s", code, size), nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, `{"error":{"message":"size should be a number"}}`, response.Body.String())
}

func TestGetNeighbors_FailServiceError(t *testing.T) {
	ctrl, mockColorService, _, _ := getMocks(t)
	defer ctrl.Finish()

	code, size, serviceError := "ff0000", 1, "error message from service"
	mockColorService.EXPECT().GetNeighbors(code, size).Return([]string{}, errors.New(serviceError))
	controller := NewColorController(mockColorService, nil, nil)

	response := getResponseRecorder("/:code", controller.GetNeighbors,
		http.MethodPost, fmt.Sprintf("/%s?size=%d", code, size), nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"error":{"message":"%s"}}`, serviceError), response.Body.String())
}

func getResponseRecorder(pathParam string, handlerFunc gin.HandlerFunc, method, query string, body io.Reader) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Any(fmt.Sprintf("/test%s", pathParam), handlerFunc)

	request, _ := http.NewRequest(method, fmt.Sprintf("/test%s", query), body)
	if method != http.MethodGet {
		request.Header.Set("Content-Type", "application/json")
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	return recorder
}

func getMocks(t *testing.T) (*gomock.Controller, *mock_services.MockColorService,
		*mock_services.MockUserService, *mock_client.MockClient) {
	ctrl := gomock.NewController(t)
	mockColorService := mock_services.NewMockColorService(ctrl)
	mockUserService := mock_services.NewMockUserService(ctrl)
	mockClient := mock_client.NewMockClient(ctrl)
	return ctrl, mockColorService, mockUserService, mockClient
}

func authorizationSuccess(user *mock_services.MockUserService, client *mock_client.MockClient) (
		*mock_services.MockUserService, *mock_client.MockClient) {
	client.EXPECT().GetUserIDFunc(gomock.Any()).Return(func() (string, error) { return "", nil })
	user.EXPECT().IsLoggedIn(gomock.Any()).Return(true)
	return user, client
}

func authorizationFail(user *mock_services.MockUserService, client *mock_client.MockClient) (
		*mock_services.MockUserService, *mock_client.MockClient) {
	client.EXPECT().GetUserIDFunc(gomock.Any()).Return(func() (string, error) { return "", nil })
	user.EXPECT().IsLoggedIn(gomock.Any()).Return(false)
	return user, client
}

func colorFormatValid(color *mock_services.MockColorService) *mock_services.MockColorService {
	color.EXPECT().IsValidCodeFormat(gomock.Any()).Return(true, "")
	return color
}

func colorFormatInvalid(color *mock_services.MockColorService) (*mock_services.MockColorService, string) {
	message := "proper regex string"
	color.EXPECT().IsValidCodeFormat(gomock.Any()).Return(false, message)
	return color, message
}

func runAdd(color *mock_services.MockColorService, client *mock_client.MockClient, lang, name, code string) (
		*mock_services.MockColorService, *mock_client.MockClient) {
	client.EXPECT().GetUserIDFunc(gomock.Any()).Return(func() (string, error) { return "", nil })
	color.EXPECT().Add(lang, name, code, gomock.Any())
	return color, client
}
