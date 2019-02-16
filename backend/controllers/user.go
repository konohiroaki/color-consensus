package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/models"
	"github.com/twinj/uuid"
	"net/http"
	"time"
)

type UserController struct{}

func (UserController) GetPresence(c *gin.Context) {
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

func (UserController) RegisterUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user) // TODO: error handling // if err := bind(); err != nil { handing... )
	user.ID = uuid.NewV4().String()
	user.Date = time.Now()
	models.Users = append(models.Users, &user)
	// TODO: move session related logic to non-api endpoint.
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()
	c.JSON(200, user);
}

func (UserController) GetUserList(c *gin.Context) {
	c.JSON(200, models.Users)
}

func findUser(userID string) (*models.User, bool) {
	for _, user := range models.Users {
		if user.ID == userID {
			return user, true
		}
	}
	return nil, false
}
