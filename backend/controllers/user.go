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
		ctx.JSON(http.StatusNotFound, errorResponse("user is not logged in"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"userID": userID})
}

func (uc UserController) Login(ctx *gin.Context) {
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("user ID should be in the request"))
		return
	}

	success := uc.userService.TryLogin(req.ID, uc.client.SetUserIDFunc(ctx))
	if !success {
		ctx.JSON(http.StatusUnauthorized, errorResponse("userID not found in repository"))
		return
	}
	ctx.Status(http.StatusOK);
}

func (uc UserController) SingUpAndLogin(ctx *gin.Context) {
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

	id, success := uc.userService.SingUpAndLogin(req.Nationality, req.Gender, req.Birth, uc.client.SetUserIDFunc(ctx))
	if !success {
		ctx.JSON(http.StatusInternalServerError, errorResponse("internal server error"))
		return
	}
	ctx.JSON(http.StatusOK, id);
}
