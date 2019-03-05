package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/domains/vote"
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
	lang := c.Param("lang")
	color := c.Param("color")
	list := vote.FindList(lang, color)
	c.JSON(200, list)
}
