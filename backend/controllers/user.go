package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"log"
	"net/http"
)

type UserController struct{}

func (UserController) GetUserIDFromCookie(ctx *gin.Context) {
	repository := ctx.Keys["userRepository"].(repositories.UserRepository)
	session := sessions.Default(ctx)
	userID := session.Get("userID")
	if userID == nil {
		ctx.Status(http.StatusNotFound)
	} else if repository.IsPresent(userID.(string)) {
		ctx.JSON(http.StatusOK, gin.H{"userID": userID})
	} else {
		// this case shouldn't exist
		ctx.Status(http.StatusInternalServerError)
	}
}

func (uc UserController) SetCookieIfUserExist(ctx *gin.Context) {
	repository := ctx.Keys["userRepository"].(repositories.UserRepository)
	type request struct {
		ID string `json:"id" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	if repository.IsPresent(req.ID) {
		if err := uc.setUserIDCookie(ctx, req.ID); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusOK);
	} else {
		log.Println("userID not found in repository")
		ctx.Status(http.StatusUnauthorized)
	}
}

func (uc UserController) AddUserAndSetCookie(ctx *gin.Context) {
	repository := ctx.Keys["userRepository"].(repositories.UserRepository)
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
	id := repository.Add(req.Nationality, req.Gender, req.Birth)
	if err := uc.setUserIDCookie(ctx, id); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, id);
}

func (UserController) setUserIDCookie(ctx *gin.Context, id string) error {
	session := sessions.Default(ctx)
	session.Set("userID", id)
	if err := session.Save(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
