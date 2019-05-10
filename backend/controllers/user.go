package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
)

type UserController struct {
	userService services.UserService
	client      client.Client
}

func NewUserController(userService services.UserService, client client.Client) UserController {
	return UserController{userService, client}
}

func (uc UserController) GetIDIfLoggedIn(ctx *gin.Context) {
	userID, err := uc.userService.GetID(uc.client.GetUserIDFunc(ctx));
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"userID": userID})
}

func (uc UserController) Login(ctx *gin.Context) {
	type request struct {
		ID string `json:"userID" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("userID should be in the request"))
		return
	}

	success := uc.userService.TryLogin(req.ID, uc.client.SetUserIDFunc(ctx))
	if !success {
		ctx.JSON(http.StatusUnauthorized, errorResponse("userID not found in repository"))
		return
	}
	ctx.Status(http.StatusOK);
}

func (uc UserController) SignUpAndLogin(ctx *gin.Context) {
	type request struct {
		Nationality string `json:"nationality" binding:"required"`
		Birth       int    `json:"birth" binding:"required"`
		Gender      string `json:"gender" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("all nationality, birth, gender should be in the request"))
		return
	}

	userID, err := uc.userService.SignUpAndLogin(req.Nationality, req.Birth, req.Gender, uc.client.SetUserIDFunc(ctx))
	if err != nil {
		switch e := err.(type) {
		case *services.ValidationError:
			ctx.JSON(http.StatusBadRequest, errorResponse(e.Error()))
			return
		case *services.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, errorResponse("internal server error"))
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"userID": userID});
}
