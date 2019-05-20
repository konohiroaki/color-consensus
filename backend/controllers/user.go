package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/services"
	"net/http"
	"sync"
)

type UserController struct {
	userService services.UserService
	client      client.Client
}

var (
	userControllerInstance UserController
	userControllerOnce     sync.Once
)

func GetUserController(env string) UserController {
	userControllerOnce.Do(func() {
		userControllerInstance = newUserController(services.GetUserService(env), client.GetClient())
	})
	return userControllerInstance
}

func newUserController(userService services.UserService, client client.Client) UserController {
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
		ID string `json:"userID" binding:"required,len=36"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(getBindErrorMessage(err)))
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
		Birth       int    `json:"birth" binding:"required,min=1900"`
		Gender      string `json:"gender" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(getBindErrorMessage(err)))
		return
	}

	userID, err := uc.userService.SignUpAndLogin(req.Nationality, req.Birth, req.Gender, uc.client.SetUserIDFunc(ctx))
	if err != nil {
		switch err.(type) {
		case *services.ValidationError:
			ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
			return
		case *services.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, errorResponse("internal server error"))
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"userID": userID});
}
