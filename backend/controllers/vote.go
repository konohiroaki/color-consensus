package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	repo "github.com/konohiroaki/color-consensus/backend/repositories"
	"log"
	"net/http"
	"strings"
)

type VoteController struct{}

func (VoteController) Vote(ctx *gin.Context) {
	voteRepo := repo.Vote(ctx)
	userID, err := client.GetUserID(ctx)
	if err != nil {
		//ctx.Status(http.StatusForbidden) // TODO: temporary skipping auth. remove this.
		//return
		userID = "testuser"
	}

	type request struct {
		Lang   string   `json:"lang" binding:"required"`
		Name   string   `json:"name" binding:"required"`
		Colors []string `json:"colors" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("all lang, name, colors should be in the request"))
		return
	}
	voteRepo.Add(userID, req.Lang, req.Name, req.Colors)
	ctx.Status(http.StatusOK)
}

func (VoteController) GetVotes(ctx *gin.Context) {
	voteRepo := repo.Vote(ctx)
	type request struct {
		Lang   string `form:"lang"`
		Name   string `form:"name"`
		Fields string `form:"fields" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("fields should be in the request"))
		return
	}
	fields := strings.Split(req.Fields, ",")
	votes := voteRepo.GetVotes(req.Lang, req.Name, fields)
	ctx.JSON(http.StatusOK, votes)
}

func (VoteController) DeleteVotesForUser(ctx *gin.Context) {
	voteRepo := repo.Vote(ctx)
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("user ID should be in the request"))
		return
	}
	voteRepo.RemoveForUser(req.ID)
	ctx.Status(http.StatusOK)
}
