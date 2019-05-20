package controllers

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/client/mock_client"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"github.com/konohiroaki/color-consensus/backend/services/mock_services"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestColorController_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService := mockColorService(ctrl)

	category, name, code := "X11", "Red", "#ff0000"
	mockColorService.EXPECT().GetAll().Return(
		[]map[string]interface{}{{"category": category, "name": name, "code": code}})
	controller := newColorController(mockColorService, nil, nil)

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`[{"category":"%s","code":"%s","name":"%s"}]`, category, code, name), response.Body.String())
}

func TestColorController_Add_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService, mockUserService, mockClient := mockColorService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	category, name, code := "X11", "Red", "#ff0000"
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	mockColorService, mockClient = doAdd(mockColorService, mockClient, category, name, code, nil)
	controller := newColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"category":"%s","name":"%s","code":"%s"}`, category, name, code))))

	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestColorController_Add_FailAuthorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService, mockUserService, mockClient := mockColorService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	mockUserService, mockClient = authorizationFail(mockUserService, mockClient)
	controller := newColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(`{"category":"X11","name":"Red","code":"#ff0000"}`)))

	assert.Equal(t, http.StatusForbidden, response.Code)
	assertErrorMessageEqual(t, "user need to be logged in to add a color", response.Body)
}

func TestColorController_Add_FailBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService, mockUserService, mockClient := mockColorService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	controller := newColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(`{"category":"X11","code":"#ff0000"}`))) // "name" not sent

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageContains(t, "Name: required", response.Body)
}

func TestColorController_Add_FailColorFormatValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService, mockUserService, mockClient := mockColorService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	controller := newColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(`{"category":"X11","name":"Red","code":"ff0000"}`)))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageContains(t, "Code: hexcolor", response.Body)
}

func TestColorController_Add_FailServiceDuplicateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService, mockUserService, mockClient := mockColorService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	category, name, code, errorMessage := "X11", "Red", "#ff0000", "error message"
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	mockColorService, mockClient = doAdd(mockColorService, mockClient, category, name, code, repositories.NewDuplicateError(errorMessage))
	controller := newColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"category":"%s","name":"%s","code":"%s"}`, category, name, code))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, errorMessage, response.Body)
}

func TestColorController_Add_FailServiceInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService, mockUserService, mockClient := mockColorService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	category, name, code, serviceError := "X11", "Red", "#ff0000", "error message"
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	mockColorService, mockClient = doAdd(mockColorService, mockClient, category, name, code, fmt.Errorf(serviceError))
	controller := newColorController(mockColorService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Add,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"category":"%s","name":"%s","code":"%s"}`, category, name, code))))

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assertErrorMessageEqual(t, serviceError, response.Body)
}

func TestColorController_GetNeighbors_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService := mockColorService(ctrl)

	code, size := "ff0000", 1
	mockColorService.EXPECT().GetNeighbors(code, size).Return([]string{"#ff0000"}, nil)
	controller := newColorController(mockColorService, nil, nil)

	response := getResponseRecorder("/:code", controller.GetNeighbors,
		http.MethodPost, fmt.Sprintf("/%s?size=%d", code, size), nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `["#ff0000"]`, response.Body.String())
}

func TestColorController_GetNeighbors_FailSizeAtoiConversion(t *testing.T) {
	code, size := "ff0000", "a"
	controller := newColorController(nil, nil, nil)

	response := getResponseRecorder("/:code", controller.GetNeighbors,
		http.MethodPost, fmt.Sprintf("/%s?size=%s", code, size), nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "size should be a number", response.Body)
}

func TestColorController_GetNeighbors_FailService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorService := mockColorService(ctrl)

	code, size, serviceError := "ff0000", 1, "error message from service"
	mockColorService.EXPECT().GetNeighbors(code, size).Return([]string{}, errors.New(serviceError))
	controller := newColorController(mockColorService, nil, nil)

	response := getResponseRecorder("/:code", controller.GetNeighbors,
		http.MethodPost, fmt.Sprintf("/%s?size=%d", code, size), nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, serviceError, response.Body)
}

func doAdd(color *mock_services.MockColorService, client *mock_client.MockClient, category, name, code string, err error) (
		*mock_services.MockColorService, *mock_client.MockClient) {
	client.EXPECT().GetUserIDFunc(gomock.Any()).Return(func() (string, error) { return "", nil })
	color.EXPECT().Add(category, name, code, gomock.Any()).Return(err)
	return color, client
}
