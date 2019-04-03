package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/domains/color"
)

type ColorController struct{}

func (ColorController) GetAll(c *gin.Context) {
	c.JSON(200, color.GetAll([]string{"lang", "name", "code"}))
}
