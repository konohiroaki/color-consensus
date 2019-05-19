package controllers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLanguageController_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLangService := mockLangService(ctrl)

	key, value := "iso_639-1", "language name in English"
	mockLangService.EXPECT().GetAll().Return(map[string]string{key: value})
	controller := newLanguageController(mockLangService)

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"%s":"%s"}`, key, value), response.Body.String())
}
