package controllers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenderController_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	genders := []interface{}{"Female", "Male", "Others"}
	controller := NewGenderController()

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`["%s","%s","%s"]`, genders...), response.Body.String())
}
