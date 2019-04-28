package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	repo "github.com/konohiroaki/color-consensus/backend/repositories"
	"log"
	"net/http"
)

type UserController struct{}

func (UserController) GetUserIDFromCookie(ctx *gin.Context) {
	userRepo := repo.User(ctx)
	userID, err := client.GetUserID(ctx)

	if err != nil {
		ctx.Status(http.StatusNotFound)
	} else if userRepo.IsPresent(userID) {
		ctx.JSON(http.StatusOK, gin.H{"userID": userID})
	} else {
		// this case shouldn't exist
		ctx.Status(http.StatusInternalServerError)
	}
}

func (UserController) SetCookieIfUserExist(ctx *gin.Context) {
	userRepo := repo.User(ctx)
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	if userRepo.IsPresent(req.ID) {
		if err := client.SetUserID(ctx, req.ID); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusOK);
	} else {
		log.Println("userID not found in repository")
		ctx.Status(http.StatusUnauthorized)
	}
}

func (UserController) AddUserAndSetCookie(ctx *gin.Context) {
	userRepo := repo.User(ctx)
	type request struct {
		Nationality string `json:"nationality" binding:"required"`
		Gender      string `json:"gender" binding:"required"`
		Birth       int    `json:"birth" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	id := userRepo.Add(req.Nationality, req.Gender, req.Birth)
	if err := client.SetUserID(ctx, id); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, id);
}
