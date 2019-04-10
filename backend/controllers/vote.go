package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"strings"
)

type VoteController struct{}

func (VoteController) Vote(ctx *gin.Context) {
	repository := ctx.Keys["voteRepository"].(repositories.VoteRepository)
	session := sessions.Default(ctx)
	userID := session.Get("userID")
	if userID == nil {
		//ctx.AbortWithStatus(http.StatusForbidden) // temporary skipping auth
		userID = "testuser"
	}

	type request struct {
		Lang   string   `json:"lang"`
		Name   string   `json:"name"`
		Colors []string `json:"colors"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
	}
	repository.Add(userID.(string), req.Lang, req.Name, req.Colors)
	ctx.Status(200)
}

func (VoteController) GetVotes(ctx *gin.Context) {
	repository := ctx.Keys["voteRepository"].(repositories.VoteRepository)
	type request struct {
		Lang   string `form:"lang"`
		Name   string `form:"name"`
		Fields string `form:"fields"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	}
	fields := strings.Split(req.Fields, ",")
	votes := repository.GetVotes(req.Lang, req.Name, fields)
	ctx.JSON(200, votes)
}

func (VoteController) DeleteVotesForUser(ctx *gin.Context) {
	repository := ctx.Keys["voteRepository"].(repositories.VoteRepository)
	type request struct {
		ID string `json:"id"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	}
	repository.RemoveForUser(req.ID)
	ctx.Status(200)
}
