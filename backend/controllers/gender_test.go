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
	mockGenderService := mockGenderService(ctrl)

	genders := []string{"foo", "bar"}
	mockGenderService.EXPECT().GetAll().Return(genders)
	controller := newGenderController(mockGenderService)

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`["%s","%s"]`, genders[0], genders[1]), response.Body.String())
}
