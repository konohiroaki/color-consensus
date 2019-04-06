package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/domains/user"
	"net/http"
)

type UserController struct{}

func (UserController) GetUserIDFromCookie(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("userID")
	if userID == nil {
		ctx.Status(http.StatusNotFound)
	} else if user.IsPresent(userID.(string)) {
		ctx.JSON(http.StatusOK, gin.H{"userID": userID})
	} else {
		// this case shouldn't exist
		ctx.Status(http.StatusPaymentRequired)
	}
}

func (UserController) SetCookieIfUserExist(ctx *gin.Context) {
	var u user.User
	if err := ctx.BindJSON(&u); err != nil {
		fmt.Println(err)
	}
	if user.IsPresent(u.ID) {
		session := sessions.Default(ctx)
		session.Set("userID", u.ID)
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusNotFound)
	}
}

func (UserController) AddUserAndSetCookie(ctx *gin.Context) {
	var u user.User
	if err := ctx.BindJSON(&u); err != nil {
		fmt.Println(err)
	}
	u = user.Add(u)
	session := sessions.Default(ctx)
	session.Set("userID", u.ID)
	if err := session.Save(); err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, u);
}
