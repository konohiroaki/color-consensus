package controllers

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/konohiroaki/color-consensus/backend/client/mock_client"
	"github.com/konohiroaki/color-consensus/backend/services/mock_services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestVoteController_Vote_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockVoteService, mockUserService, mockClient := mockVoteService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	category, name, colors := "X11", "Red", []string{"#000000", "#000010"}
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	mockVoteService, mockClient = doVote(mockVoteService, mockClient, category, name, colors)
	controller := newVoteController(mockVoteService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Vote,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"category":"%s","name":"%s","colors":["%s"]}`, category, name, strings.Join(colors, "\",\"")))))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestVoteController_Vote_FailAuthorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockVoteService, mockUserService, mockClient := mockVoteService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	category, name, colors := "X11", "Red", []string{"#000000", "#000010"}
	mockUserService, mockClient = authorizationFail(mockUserService, mockClient)
	controller := newVoteController(mockVoteService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Vote,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"category":"%s","name":"%s","colors":["%s"]}`, category, name, strings.Join(colors, "\",\"")))))

	assert.Equal(t, http.StatusForbidden, response.Code)
	assertErrorMessageEqual(t, "user need to be logged in to vote", response.Body)
}

func TestVoteController_Vote_FailBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockVoteService, mockUserService, mockClient := mockVoteService(ctrl), mockUserService(ctrl), mockClient(ctrl)

	category, name := "X11", "Red"
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	controller := newVoteController(mockVoteService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Vote,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"category":"%s","name":"%s"}`, category, name))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "all category, name, colors should be in the request", response.Body)
}

func TestVoteController_Get_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockVoteService := mockVoteService(ctrl)

	category, name, fields := "X11", "Red", []string{"category"}
	mockVoteService.EXPECT().Get(category, name, fields).Return([]map[string]interface{}{{"category": "X11"}})
	controller := newVoteController(mockVoteService, nil, nil)

	response := getResponseRecorder("", controller.Get,
		http.MethodGet, fmt.Sprintf("?category=%s&name=%s&fields=%s", category, name, strings.Join(fields, ",")), nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `[{"category":"X11"}]`, response.Body.String())
}

func TestVoteController_Get_FailBind(t *testing.T) {
	category, name := "X11", "Red"
	controller := newVoteController(nil, nil, nil)

	response := getResponseRecorder("", controller.Get,
		http.MethodGet, fmt.Sprintf("?category=%s&name=%s", category, name), nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "fields should be in the request", response.Body)
}

func TestVoteController_RemoveByUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockVoteService := mockVoteService(ctrl)

	userID := "id"
	mockVoteService.EXPECT().RemoveByUser(userID)
	controller := newVoteController(mockVoteService, nil, nil)

	response := getResponseRecorder("", controller.RemoveByUser,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, userID))))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestVoteController_RemoveByUser_FailBind(t *testing.T) {
	userID := "id"
	controller := newVoteController(nil, nil, nil)

	response := getResponseRecorder("", controller.RemoveByUser,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"user":"%s"}`, userID))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "userID should be in the request", response.Body)
}

func doVote(vote *mock_services.MockVoteService, client *mock_client.MockClient, category, name string, colors []string) (
		*mock_services.MockVoteService, *mock_client.MockClient) {
	client.EXPECT().GetUserIDFunc(gomock.Any()).Return(func() (string, error) { return "", nil })
	vote.EXPECT().Vote(category, name, colors, gomock.Any())
	return vote, client
}
