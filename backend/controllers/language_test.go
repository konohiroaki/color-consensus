package controllers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLanguageController_GetAll_Success(t *testing.T) {
	ctrl, _, _, _, mockLangService, _ := getMocks(t)
	defer ctrl.Finish()

	key, value := "iso_639-1", "language name in English"
	mockLangService.EXPECT().GetAll().Return(map[string]string{key: value})
	controller := NewLanguageController(mockLangService)

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`{"%s":"%s"}`, key, value), response.Body.String())
}
