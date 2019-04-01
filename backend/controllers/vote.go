package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/domains/user"
	"github.com/konohiroaki/color-consensus/backend/domains/vote"
	"strings"
)

type VoteController struct{}

func (VoteController) Vote(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		//c.AbortWithStatus(http.StatusForbidden) // temporary skipping auth
		userID = "testuser"
	}

	type request struct {
		Lang   string `json:"lang"`
		Name   string `json:"name"`
		Colors []string `json:"colors"`
	}
	var req request
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
	}
	vote.Add(userID.(string), req.Lang, req.Name, req.Colors)
	c.Status(200)
}

func (VoteController) GetVotes(c *gin.Context) {
	type request struct {
		Lang   string `form:"lang"`
		Name   string `form:"name"`
		Fields string `form:"fields"`
	}
	var req request
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	fields := strings.Split(req.Fields, ",")
	votes := vote.GetVotes(req.Lang, req.Name, fields)
	c.JSON(200, votes)
}

func (VoteController) DeleteVotesForUser(c *gin.Context) {
	var user user.User
	_ = c.BindJSON(&user)
	vote.RemoveForUser(user.ID)
	c.Status(200)
}
