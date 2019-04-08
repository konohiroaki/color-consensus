package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"net/http"
	"strings"
)

type VoteController struct {
}

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

func (VoteController) GetStats(ctx *gin.Context) {
	repository := ctx.Keys["voteRepository"].(repositories.VoteRepository)
	type request struct {
		Lang string `form:"lang"`
		Name string `form:"name"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(400)
		return
	}
	votes := repository.GetVotes(req.Lang, req.Name, []string{"colors"})
	type response struct {
		Count int            `json:"count"`
		Codes map[string]int `json:"colors"`
	}
	res := response{}
	res.Count = len(votes)
	res.Codes = map[string]int{}
	for _, vote := range votes {
		codes := vote["colors"].([]interface{})
		for _, code := range codes {
			c := code.(string)
			if val, ok := res.Codes[c]; ok {
				res.Codes[c] = val + 1
			} else {
				res.Codes[c] = 1
			}
		}
	}
	ctx.JSON(http.StatusOK, res)
}

func (VoteController) DeleteVotesForUser(ctx *gin.Context) {
	repository := ctx.Keys["voteRepository"].(repositories.VoteRepository)
	type request struct {
		ID string `json:"id"`
	}
	var req request
	_ = ctx.BindJSON(&req)
	repository.RemoveForUser(req.ID)
	ctx.Status(200)
}
