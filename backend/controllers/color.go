package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
	"strconv"
)

type ColorController struct{}

func NewColorController() ColorController {
	return ColorController{}
}

func (ColorController) GetAll(ctx *gin.Context) {
	colors := services.Color(ctx).GetAll(ctx)

	ctx.JSON(http.StatusOK, colors)
	return
}

func (ColorController) Add(ctx *gin.Context) {
	if !services.User(ctx).IsLoggedIn(ctx) {
		ctx.JSON(http.StatusForbidden, errorResponse("user need to be logged in to add a color"))
		return
	}

	type request struct {
		Lang string `json:"lang" binding:"required"`
		Name string `json:"name" binding:"required"`
		Code string `json:"code" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("all language, name, code are necessary"))
		return
	}

	if ok, regex := services.Color(ctx).IsValidCodeFormat(req.Code); !ok {
		ctx.JSON(http.StatusBadRequest, errorResponse("color code should match regex: "+regex))
		return
	}

	services.Color(ctx).Add(ctx, req.Lang, req.Name, req.Code)
	ctx.Status(http.StatusCreated);
}

func (ColorController) GetNeighbors(ctx *gin.Context) {
	code := ctx.Param("code")
	size, sizeErr := strconv.Atoi(ctx.Query("size"));
	if sizeErr != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("size should be a number"))
		return
	}

	neighbors, serviceErr := services.Color(ctx).GetNeighbors(code, size)
	if serviceErr != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(serviceErr.Error()))
		return
	}

	ctx.JSON(http.StatusOK, neighbors)
}

func errorResponse(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"message": message,
		},
	}
}
