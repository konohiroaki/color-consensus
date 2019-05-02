package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
	"strconv"
)

type ColorController struct {
	colorService services.ColorService
	userService  services.UserService
}

func NewColorController(colorService services.ColorService, userService services.UserService) ColorController {
	return ColorController{colorService, userService}
}

func (cc ColorController) GetAll(ctx *gin.Context) {
	colors := cc.colorService.GetAll()

	ctx.JSON(http.StatusOK, colors)
	return
}

func (cc ColorController) Add(ctx *gin.Context) {
	if !cc.userService.IsLoggedIn(client.GetUserIDFunc(ctx)) {
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

	if ok, regex := cc.colorService.IsValidCodeFormat(req.Code); !ok {
		ctx.JSON(http.StatusBadRequest, errorResponse("color code should match regex: "+regex))
		return
	}

	userID, _ := client.GetUserIDFunc(ctx)()
	cc.colorService.Add(req.Lang, req.Name, req.Code, userID)
	ctx.Status(http.StatusCreated);
}

func (cc ColorController) GetNeighbors(ctx *gin.Context) {
	code := ctx.Param("code")
	size, sizeErr := strconv.Atoi(ctx.Query("size"));
	if sizeErr != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("size should be a number"))
		return
	}

	neighbors, serviceErr := cc.colorService.GetNeighbors(code, size)
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
