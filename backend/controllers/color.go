package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/models"
)

type ColorController struct{}

func (ColorController) GetAllConsensus(c *gin.Context) {
	//TODO: provide a way to get only language list
	c.JSON(200, models.Sum)
}

func (ColorController) GetAllConsensusForLang(c *gin.Context) {
	list := findSumListOfLang(c.Param("lang"))
	c.JSON(200, list)
}

func (ColorController) GetConsensus(c *gin.Context) {
	lang := c.Param("lang")
	color := c.Param("color")
	sum, found := findSum(lang, color)
	if found {
		c.JSON(200, sum)
	} else {
		c.AbortWithStatus(404)
	}
}

func findSumListOfLang(lang string) []models.ColorConsensus {
	list := []models.ColorConsensus{}
	for _, ele := range models.Sum {
		if ele.Language == lang {
			list = append(list, *ele)
		}
	}
	return list
}
