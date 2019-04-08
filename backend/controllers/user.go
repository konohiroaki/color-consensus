package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"net/http"
)

type UserController struct{}

func (UserController) GetUserIDFromCookie(ctx *gin.Context) {
	repository := ctx.Keys["userRepository"].(repositories.UserRepository)
	session := sessions.Default(ctx)
	userID := session.Get("userID")
	if userID == nil {
		ctx.Status(http.StatusNotFound)
	} else if repository.IsPresent(userID.(string)) {
		ctx.JSON(http.StatusOK, gin.H{"userID": userID})
	} else {
		// this case shouldn't exist
		ctx.Status(http.StatusPaymentRequired)
	}
}

func (UserController) SetCookieIfUserExist(ctx *gin.Context) {
	repository := ctx.Keys["userRepository"].(repositories.UserRepository)
	type request struct {
		ID string `json:"id"`
	}
	var req request
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println(err)
	}
	if repository.IsPresent(req.ID) {
		session := sessions.Default(ctx)
		session.Set("userID", req.ID)
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusNotFound)
	}
}

func (UserController) AddUserAndSetCookie(ctx *gin.Context) {
	repository := ctx.Keys["userRepository"].(repositories.UserRepository)
	type request struct {
		Nationality string `json:"nationality"`
		Gender      string `json:"gender"`
		Birth       int    `json:"birth"`
	}
	var req request
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println(err)
	}
	id := repository.Add(req.Nationality, req.Gender, req.Birth)
	session := sessions.Default(ctx)
	session.Set("userID", id)
	if err := session.Save(); err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, id);
}
