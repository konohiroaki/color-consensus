package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/models"
)

type ColorController struct{}

func (ColorController) GetAllConsensusKey(c *gin.Context) {
	type ResponseElement struct {
		Language string
		Color    string
	}
	list := []ResponseElement{}
	for _, e := range models.Consensus {
		list = append(list, ResponseElement{e.Language, e.Color})
	}
	c.JSON(200, list)
}

func (ColorController) GetAllConsensus(c *gin.Context) {
	c.JSON(200, models.Consensus)
}

func (ColorController) GetAllConsensusKeyForLang(c *gin.Context) {
	lang := c.Param("lang")
	list := []string{}
	for _, e := range models.Consensus {
		if e.Language == lang {
			list = append(list, e.Color)
		}
	}
	c.JSON(200, list)
}

func (ColorController) GetAllConsensusForLang(c *gin.Context) {
	list := findConsensusListOfLang(c.Param("lang"))
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

func findConsensusListOfLang(lang string) []models.ColorConsensus {
	list := []models.ColorConsensus{}
	for _, ele := range models.Consensus {
		if ele.Language == lang {
			list = append(list, *ele)
		}
	}
	return list
}
