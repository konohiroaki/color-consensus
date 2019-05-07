package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/client/mock_client"
	"github.com/konohiroaki/color-consensus/backend/services/mock_services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_ErrorResponse(t *testing.T) {
	actual := errorResponse("message")

	byteArr, _ := json.Marshal(actual)
	jsonStr := string(byteArr)

	assert.Equal(t, `{"error":{"message":"message"}}`, jsonStr)
}

func assertErrorMessageEqual(t *testing.T, expected string, actual *bytes.Buffer) {
	assert.Equal(t, fmt.Sprintf(`{"error":{"message":"%s"}}`, expected), actual.String())
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

func mockColorService(ctrl *gomock.Controller) *mock_services.MockColorService {
	return mock_services.NewMockColorService(ctrl)
}

func mockVoteService(ctrl *gomock.Controller) *mock_services.MockVoteService {
	return mock_services.NewMockVoteService(ctrl)
}

func mockUserService(ctrl *gomock.Controller) *mock_services.MockUserService {
	return mock_services.NewMockUserService(ctrl)
}

func mockLangService(ctrl *gomock.Controller) *mock_services.MockLanguageService {
	return mock_services.NewMockLanguageService(ctrl)
}

func mockGenderService(ctrl *gomock.Controller) *mock_services.MockGenderService {
	return mock_services.NewMockGenderService(ctrl)
}

func mockClient(ctrl *gomock.Controller) *mock_client.MockClient {
	return mock_client.NewMockClient(ctrl)
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
