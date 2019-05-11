package controllers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNationalityController_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockNationService := mockNationService(ctrl)

	key, value := "iso_3166-1", "nationality name in English"
	mockNationService.EXPECT().GetAll().Return(map[string]string{key: value})
	controller := NewNationalityController(mockNationService)

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"%s":"%s"}`, key, value), response.Body.String())
}
