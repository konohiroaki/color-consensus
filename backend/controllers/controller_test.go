package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/client/mock_client"
	"github.com/konohiroaki/color-consensus/backend/services/mock_services"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func getMocks(t *testing.T) (*gomock.Controller, *mock_services.MockColorService, *mock_services.MockVoteService,
		*mock_services.MockUserService, *mock_services.MockLanguageService, *mock_client.MockClient) {
	ctrl := gomock.NewController(t)
	mockColorService := mock_services.NewMockColorService(ctrl)
	mockVoteService := mock_services.NewMockVoteService(ctrl)
	mockUserService := mock_services.NewMockUserService(ctrl)
	mockLangService := mock_services.NewMockLanguageService(ctrl)
	mockClient := mock_client.NewMockClient(ctrl)
	return ctrl, mockColorService, mockVoteService, mockUserService, mockLangService, mockClient
}
