package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
	"strings"
	"sync"
)

type VoteController struct {
	voteService services.VoteService
	userService services.UserService
	client      client.Client
}

var (
	voteControllerInstance VoteController
	voteControllerOnce     sync.Once
)

func GetVoteController(env string) VoteController {
	voteControllerOnce.Do(func() {
		voteControllerInstance = newVoteController(services.GetVoteService(env), services.GetUserService(env), client.GetClient())
	})
	return voteControllerInstance
}

func newVoteController(voteService services.VoteService, userService services.UserService, client client.Client) VoteController {
	return VoteController{voteService, userService, client}
}

func (vc VoteController) Vote(ctx *gin.Context) {
	if !vc.userService.IsLoggedIn(vc.client.GetUserIDFunc(ctx)) {
		ctx.JSON(http.StatusForbidden, errorResponse("user need to be logged in to vote"))
		return
	}

	type request struct {
		Category   string   `json:"category" binding:"required"`
		Name   string   `json:"name" binding:"required"`
		Colors []string `json:"colors" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("all category, name, colors should be in the request"))
		return
	}

	vc.voteService.Vote(req.Category, req.Name, req.Colors, vc.client.GetUserIDFunc(ctx))
	ctx.Status(http.StatusOK)
}

func (vc VoteController) Get(ctx *gin.Context) {
	type request struct {
		Category   string `form:"category"`
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

	votes := vc.voteService.Get(req.Category, req.Name, fields)
	ctx.JSON(http.StatusOK, votes)
}

func (vc VoteController) RemoveByUser(ctx *gin.Context) {
	type request struct {
		ID string `json:"userID" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("userID should be in the request"))
		return
	}

	vc.voteService.RemoveByUser(req.ID)
	ctx.Status(http.StatusOK)
}
