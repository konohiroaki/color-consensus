package controllers

import (
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

	var v vote.ColorVote
	_ = c.BindJSON(&v)
	v.User = userID.(string)

	vote.Add(v)

	c.Status(200)
}

func (VoteController) GetVotes(c *gin.Context) {
	type request struct {
		Lang   string `form:"lang"`
		Name   string `form:"name"`
		Fields string `form:"fields"`
	}
	var req request
	if (c.ShouldBind(&req)) == nil {
		fields := strings.Split(req.Fields, ",")
		votes := vote.GetVotes(req.Lang, req.Name, fields)
		c.JSON(200, votes)
		return
	}
	c.AbortWithStatus(400)
}

func (VoteController) DeleteVotesForUser(c *gin.Context) {
	var user user.User
	_ = c.BindJSON(&user)
	vote.RemoveForUser(user.ID)
	c.Status(200)
}
