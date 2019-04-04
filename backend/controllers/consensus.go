package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/domains/consensus"
	"net/http"
)

type ConsensusController struct{}

func (ConsensusController) GetAllConsensusKey(c *gin.Context) {
	c.JSON(200, consensus.GetKeys())
}

func (ConsensusController) GetAllConsensus(c *gin.Context) {
	c.JSON(200, consensus.GetList())
}

func (ConsensusController) GetConsensus(c *gin.Context) {
	lang := c.Param("lang")
	color := c.Param("color")
	cc, found := consensus.Get(lang, color)
	if found {
		c.JSON(200, cc)
	} else {
		c.AbortWithStatus(404)
	}
}

func (ConsensusController) AddColor(c *gin.Context) {
	var colorConsensus consensus.ColorConsensus
	_ = c.BindJSON(&colorConsensus)
	consensus.Add(colorConsensus)

	c.Status(http.StatusCreated);
}
