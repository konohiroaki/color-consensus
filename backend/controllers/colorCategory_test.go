package controllers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestColorCategoryController_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockColorCategoryService := mockColorCategoryService(ctrl)

	categories := []string{"a", "b", "c"}
	mockColorCategoryService.EXPECT().GetAll().Return(categories)
	controller := newColorCategoryController(mockColorCategoryService)

	response := getResponseRecorder("", controller.GetAll, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, fmt.Sprintf(`["%s","%s","%s"]`, categories[0], categories[1], categories[2]), response.Body.String())
}
