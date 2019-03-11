package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/domains/user"
	"net/http"
)

type UserController struct{}

func (UserController) GetUserIDFromCookie(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.Status(http.StatusNotFound)
	} else if _, found := user.Get(userID.(string)); found {
		c.JSON(http.StatusOK, gin.H{"userID": userID})
	} else {
		// this case shouldn't exist
		c.Status(http.StatusPaymentRequired)
	}
}

func (UserController) SetCookieIfUserExist(c *gin.Context) {
	var u user.User
	if err := c.BindJSON(&u); err != nil {
		fmt.Println(err)
	}
	if _, found := user.Get(u.ID); found {
		session := sessions.Default(c)
		session.Set("userID", u.ID)
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (UserController) AddUserAndSetCookie(c *gin.Context) {
	var u user.User
	if err := c.BindJSON(&u); err != nil {
		fmt.Println(err)
	}
	u = user.Add(u)
	session := sessions.Default(c)
	session.Set("userID", u.ID)
	if err := session.Save(); err != nil {
		fmt.Println(err)
	}
	c.JSON(200, u);
}
