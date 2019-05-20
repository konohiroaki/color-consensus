package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/services"
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
		Category string   `json:"category" binding:"required,max=20"`
		Name     string   `json:"name" binding:"required,max=30"`
		Colors   []string `json:"colors" binding:"required,dive,hexcolor"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(getBindErrorMessage(err)))
		return
	}

	vc.voteService.Vote(req.Category, req.Name, req.Colors, vc.client.GetUserIDFunc(ctx))
	ctx.Status(http.StatusOK)
}

func (vc VoteController) Get(ctx *gin.Context) {
	type request struct {
		Category string `form:"category" binding:"max=20"`
		Name     string `form:"name" binding:"max=30"`
		Fields   string `form:"fields" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(getBindErrorMessage(err)))
		return
	}
	fields := strings.Split(req.Fields, ",")

	votes := vc.voteService.Get(req.Category, req.Name, fields)
	ctx.JSON(http.StatusOK, votes)
}

func (vc VoteController) RemoveByUser(ctx *gin.Context) {
	type request struct {
		ID string `json:"userID" binding:"required,len=36"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(getBindErrorMessage(err)))
		return
	}

	vc.voteService.RemoveByUser(req.ID)
	ctx.Status(http.StatusOK)
}
