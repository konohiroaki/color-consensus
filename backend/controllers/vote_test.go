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
	ctrl, _, mockVoteService, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	lang, name, colors := "en", "red", []string{"#000000", "#000010"}
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	mockVoteService, mockClient = doVote(mockVoteService, mockClient, lang, name, colors)
	controller := NewVoteController(mockVoteService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Vote,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"lang":"%s","name":"%s","colors":["%s"]}`, lang, name, strings.Join(colors, "\",\"")))))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestVoteController_Vote_FailAuthorization(t *testing.T) {
	ctrl, _, mockVoteService, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	lang, name, colors := "en", "red", []string{"#000000", "#000010"}
	mockUserService, mockClient = authorizationFail(mockUserService, mockClient)
	controller := NewVoteController(mockVoteService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Vote,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"lang":"%s","name":"%s","colors":["%s"]}`, lang, name, strings.Join(colors, "\",\"")))))

	assert.Equal(t, http.StatusForbidden, response.Code)
	assertErrorMessageEqual(t, "user need to be logged in to vote", response.Body)
}

func TestVoteController_Vote_FailBind(t *testing.T) {
	ctrl, _, mockVoteService, mockUserService, _, mockClient := getMocks(t)
	defer ctrl.Finish()

	lang, name := "en", "red"
	mockUserService, mockClient = authorizationSuccess(mockUserService, mockClient)
	controller := NewVoteController(mockVoteService, mockUserService, mockClient)

	response := getResponseRecorder("", controller.Vote,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(
			`{"lang":"%s","name":"%s"}`, lang, name))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "all lang, name, colors should be in the request", response.Body)
}

func TestVoteController_Get_Success(t *testing.T) {
	ctrl, _, mockVoteService, _, _, _ := getMocks(t)
	defer ctrl.Finish()

	lang, name, fields := "en", "red", []string{"lang"}
	mockVoteService.EXPECT().Get(lang, name, fields).Return([]map[string]interface{}{{"lang": "en"}})
	controller := NewVoteController(mockVoteService, nil, nil)

	response := getResponseRecorder("", controller.Get,
		http.MethodGet, fmt.Sprintf("?lang=%s&name=%s&fields=%s", lang, name, strings.Join(fields, ",")), nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `[{"lang":"en"}]`, response.Body.String())
}

func TestVoteController_Get_FailBind(t *testing.T) {
	lang, name := "en", "red"
	controller := NewVoteController(nil, nil, nil)

	response := getResponseRecorder("", controller.Get,
		http.MethodGet, fmt.Sprintf("?lang=%s&name=%s", lang, name), nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "fields should be in the request", response.Body)
}

func TestVoteController_RemoveByUser_Success(t *testing.T) {
	ctrl, _, mockVoteService, _, _, _ := getMocks(t)
	defer ctrl.Finish()

	id := "id"
	mockVoteService.EXPECT().RemoveByUser(id)
	controller := NewVoteController(mockVoteService, nil, nil)

	response := getResponseRecorder("", controller.RemoveByUser,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"userID":"%s"}`, id))))

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestVoteController_RemoveByUser_FailBind(t *testing.T) {
	id := "id"
	controller := NewVoteController(nil, nil, nil)

	response := getResponseRecorder("", controller.RemoveByUser,
		http.MethodPost, "", bytes.NewBuffer([]byte(fmt.Sprintf(`{"user":"%s"}`, id))))

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assertErrorMessageEqual(t, "userID should be in the request", response.Body)
}

func doVote(vote *mock_services.MockVoteService, client *mock_client.MockClient, lang, name string, colors []string) (
		*mock_services.MockVoteService, *mock_client.MockClient) {
	client.EXPECT().GetUserIDFunc(gomock.Any()).Return(func() (string, error) { return "", nil })
	vote.EXPECT().Vote(lang, name, colors, gomock.Any())
	return vote, client
}
