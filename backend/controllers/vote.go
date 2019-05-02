package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
	"strings"
)

type VoteController struct {
	voteService services.VoteService
	userService services.UserService
}

func NewVoteController(voteService services.VoteService, userService services.UserService) VoteController {
	return VoteController{voteService, userService}
}

func (vc VoteController) Vote(ctx *gin.Context) {
	if !vc.userService.IsLoggedIn(client.GetUserIDFunc(ctx)) {
		ctx.Status(http.StatusForbidden)
		return
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

	vc.voteService.Vote(req.Lang, req.Name, req.Colors, client.GetUserIDFunc(ctx))
	ctx.Status(http.StatusOK)
}

func (vc VoteController) Get(ctx *gin.Context) {
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

	votes := vc.voteService.Get(req.Lang, req.Name, fields)
	ctx.JSON(http.StatusOK, votes)
}

func (vc VoteController) RemoveByUser(ctx *gin.Context) {
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("user ID should be in the request"))
		return
	}

	vc.voteService.RemoveByUser(req.ID)
	ctx.Status(http.StatusOK)
}
