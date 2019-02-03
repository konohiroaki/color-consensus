package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/models"
	"time"
)

type VoteController struct{}

func (VoteController) Vote(c *gin.Context) {
	// TODO: move session related logic to non-api endpoint.
	session := sessions.Default(c)
	userID := session.Get("userID")
	fmt.Println("userID is {}", session.Get("userID"))
	if userID == nil {
		//c.AbortWithStatus(http.StatusForbidden) // temporary skipping auth
		userID = "testuser"
	}

	var colors []string
	c.BindJSON(&colors) // TODO: error handling // if err := bind(); err != nil { handing... )
	lang := c.Param("lang")
	color := c.Param("color")
	vote := models.ColorVote{Language: lang, Color: color, User: userID.(string), Date: time.Now(), Colors: colors}

	models.Votes = append(models.Votes, &vote)
	sum, found := findSum(lang, color)
	if found {
		sum.Vote += 1
		// TODO: distinct color list
		for _, votedColor := range vote.Colors {
			if _, found := sum.Colors[votedColor]; found {
				sum.Colors[votedColor] += 1
			} else {
				sum.Colors[votedColor] = 1
			}
		}
		c.JSON(200, sum)
	} else {
		c.AbortWithStatus(404)
	}
}

func (VoteController) GetVotes(c *gin.Context) {
	lang := c.Param("lang")
	color := c.Param("color")
	list := findVoteList(lang, color)
	c.JSON(200, list)
}

func findSum(lang, color string) (*models.ColorConsensus, bool) {
	for _, ele := range models.Consensus {
		if ele.Language == lang && ele.Color == color {
			return ele, true
		}
	}
	return nil, false
}

func findVoteList(lang, color string) []models.ColorVote {
	list := []models.ColorVote{};
	for _, ele := range models.Votes {
		if ele.Language == lang && ele.Color == color {
			list = append(list, *ele)
		}
	}
	return list
}
