package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type colorServiceMock struct{}

func (colorServiceMock) GetAll() []map[string]interface{} {
	return []map[string]interface{}{
		{"lang": "en", "name": "red", "code": "#ff0000"},
	}
}
func (colorServiceMock) Add(lang, name, code string, getUserID func() (string, error)) {

}
func (colorServiceMock) GetNeighbors(code string, size int) ([]string, error) {
	return []string{}, nil
}
func (colorServiceMock) IsValidCodeFormat(input string) (bool, string) {
	return true, ""
}

type userServiceMock struct{}

func (userServiceMock) IsLoggedIn(getUserID func() (string, error)) bool {
	return true
}
func (userServiceMock) GetID(getUserID func() (string, error)) (string, error) {
	return "id", nil
}
func (userServiceMock) SingUpAndLogin(nationality, gender string, birth int, setUserID func(string) error) (string, bool) {
	return "newid", true
}
func (userServiceMock) TryLogin(userID string, setUserID func(string) error) bool {
	return true
}

type clientMock struct{}

func (clientMock) GetUserIDFunc(ctx *gin.Context) (func() (string, error)) {
	return func() (s string, e error) { return "", nil }
}
func (clientMock) SetUserIDFunc(ctx *gin.Context) (func(string) error) {
	return func(s string) error { return nil }
}

func TestGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := NewColorController(colorServiceMock{}, nil, nil)
	router.GET("/test", controller.GetAll)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, `[{"code":"#ff0000","lang":"en","name":"red"}]`, recorder.Body.String())
}

func TestAdd(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := NewColorController(colorServiceMock{}, userServiceMock{}, clientMock{})
	router.POST("/test", controller.Add)

	req, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(`{"code":"#ff0000","lang":"en","name":"red"}`)))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}
