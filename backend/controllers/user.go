package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (UserController) GetIDIfLoggedIn(ctx *gin.Context) {
	userID, err := services.User(ctx).GetID(ctx);
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse("user is not logged in"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"userID": userID})
}

func (UserController) Login(ctx *gin.Context) {
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("user ID should be in the request"))
		return
	}

	success := services.User(ctx).TryLogin(ctx, req.ID)
	if !success {
		ctx.JSON(http.StatusUnauthorized, errorResponse("userID not found in repository"))
		return
	}
	ctx.Status(http.StatusOK);
}

func (UserController) SingUpAndLogin(ctx *gin.Context) {
	type request struct {
		Nationality string `json:"nationality" binding:"required"`
		Gender      string `json:"gender" binding:"required"`
		Birth       int    `json:"birth" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("all nationality, gender, birth should be in the request"))
		return
	}

	id, success := services.User(ctx).SingUpAndLogin(ctx, req.Nationality, req.Gender, req.Birth)
	if !success {
		ctx.JSON(http.StatusInternalServerError, errorResponse("internal server error"))
		return
	}
	ctx.JSON(http.StatusOK, id);
}
