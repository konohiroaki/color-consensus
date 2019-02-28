package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/models"
	"github.com/konohiroaki/color-consensus/backend/repository"
	"github.com/twinj/uuid"
	"net/http"
	"time"
)

type UserController struct{}

func (UserController) GetPresenceFromCookie(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.Status(http.StatusNotFound)
	} else if _, found := findUser(userID.(string)); !found {
		// this case shouldn't exist
		c.Status(http.StatusPaymentRequired)
	} else {
		c.JSON(http.StatusOK, gin.H{"userID": userID})
	}
}

func (UserController) ConfirmPresence(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
	}
	if _, found := findUser(user.ID); !found {
		c.Status(http.StatusNotFound)
	} else {
		session := sessions.Default(c)
		session.Set("userID", user.ID)
		c.JSON(http.StatusOK, gin.H{"userID": user.ID})
	}
}

func (UserController) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		// TODO: error handling
		fmt.Println(err)
	}
	user.ID = uuid.NewV4().String()
	user.Date = time.Now()
	repository.Users = append(repository.Users, &user)
	// TODO: move session related logic to non-api endpoint.
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	if err := session.Save(); err != nil {
		// TODO: error handling
		fmt.Println(err)
	}
	c.JSON(200, user);
}

func (UserController) GetUserList(c *gin.Context) {
	c.JSON(200, repository.Users)
}

func findUser(userID string) (*models.User, bool) {
	for _, user := range repository.Users {
		if user.ID == userID {
			return user, true
		}
	}
	return nil, false
}
