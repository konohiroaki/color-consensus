package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
	"strings"
)

type VoteController struct{}

func NewVoteController() VoteController {
	return VoteController{}
}

func (VoteController) Vote(ctx *gin.Context) {
	if !services.User(ctx).IsLoggedIn(ctx) {
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

	services.Vote(ctx).Vote(ctx, req.Lang, req.Name, req.Colors)
	ctx.Status(http.StatusOK)
}

func (VoteController) Get(ctx *gin.Context) {
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

	votes := services.Vote(ctx).Get(ctx, req.Lang, req.Name, fields)
	ctx.JSON(http.StatusOK, votes)
}

func (VoteController) RemoveByUser(ctx *gin.Context) {
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("user ID should be in the request"))
		return
	}

	services.Vote(ctx).RemoveByUser(ctx, req.ID)
	ctx.Status(http.StatusOK)
}
