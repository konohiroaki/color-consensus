package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"github.com/konohiroaki/color-consensus/backend/services"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type ColorController struct {
	colorService services.ColorService
	userService  services.UserService
	client       client.Client
}

var (
	colorControllerInstance ColorController
	colorControllerOnce     sync.Once
)

func GetColorController(env string) ColorController {
	colorControllerOnce.Do(func() {
		colorControllerInstance = newColorController(services.GetColorService(env), services.GetUserService(env), client.GetClient())
	})
	return colorControllerInstance
}

func newColorController(colorService services.ColorService, userService services.UserService, client client.Client) ColorController {
	return ColorController{colorService, userService, client}
}

func (cc ColorController) GetAll(ctx *gin.Context) {
	colors := cc.colorService.GetAll()

	ctx.JSON(http.StatusOK, colors)
	return
}

func (cc ColorController) Add(ctx *gin.Context) {
	if !cc.userService.IsLoggedIn(cc.client.GetUserIDFunc(ctx)) {
		ctx.JSON(http.StatusForbidden, errorResponse("user need to be logged in to add a color"))
		return
	}

	type request struct {
		Category string `json:"category" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Code     string `json:"code" binding:"required"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse("all category, name, code are necessary"))
		return
	}

	if ok, regex := cc.colorService.IsValidCodeFormat(req.Code); !ok {
		ctx.JSON(http.StatusBadRequest, errorResponse("color code should match regex: "+regex))
		return
	}

	err := cc.colorService.Add(req.Category, req.Name, req.Code, cc.client.GetUserIDFunc(ctx))
	if err != nil {
		switch err.(type) {
		case *services.ValidationError:
			ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
			return
		case *repositories.DuplicateError:
			ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
			return
		}
	}
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
