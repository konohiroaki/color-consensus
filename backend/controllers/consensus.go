package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/models"
	"strconv"
)

type ConsensusController struct{}

func (ConsensusController) GetAllConsensusKey(c *gin.Context) {
	type ResponseElement struct {
		Language string `json:"lang"`
		Color    string `json:"name"`
		BaseCode string `json:"code"`
	}
	list := []ResponseElement{}
	for _, e := range models.Consensus {
		list = append(list, ResponseElement{e.Language, e.Color, e.BaseCode})
	}
	c.JSON(200, list)
}

func (ConsensusController) GetAllConsensus(c *gin.Context) {
	c.JSON(200, models.Consensus)
}

func (ConsensusController) GetAllConsensusKeyForLang(c *gin.Context) {
	lang := c.Param("lang")
	type ResponseElement struct {
		Language string `json:"lang"`
		Color    string `json:"name"`
		BaseCode string `json:"code"`
	}
	list := []ResponseElement{}
	for _, e := range models.Consensus {
		if e.Language == lang {
			list = append(list, ResponseElement{e.Language, e.Color, e.BaseCode})
		}
	}
	c.JSON(200, list)
}

func (ConsensusController) GetAllConsensusForLang(c *gin.Context) {
	list := findConsensusListOfLang(c.Param("lang"))
	c.JSON(200, list)
}

func (ConsensusController) GetConsensus(c *gin.Context) {
	lang := c.Param("lang")
	color := c.Param("color")
	sum, found := findSum(lang, color)
	if found {
		c.JSON(200, sum)
	} else {
		c.AbortWithStatus(404)
	}
}

func (ConsensusController) GetCandidateList(c *gin.Context) {
	code := c.Param("code")
	candidates := generateCandidateList(code)
	c.JSON(200, candidates)
}

func generateCandidateList(code string) []string {
	r := fromHex(code[0:2])
	g := fromHex(code[2:4])
	b := fromHex(code[4:])
	list := []string{}
	for x := 1; x < 100; x++ {
		for i, rr := 0, r; i <= x; i, rr = i+1, swingIncrement(r, i+1, 16) {
			for j, gg := 0, g; j <= x; j, gg = j+1, swingIncrement(g, j+1, 16) {
				for k, bb := 0, b; k <= x; k, bb = k+1, swingIncrement(b, k+1, 16) {
					if i < x && j < x && k < x {
						continue
					}
					str := "#" + toHex(rr) + toHex(gg) + toHex(bb)
					list = append(list, str)
					if len(list) == 51*51 {
						return list
					}
				}
			}
		}
	}
	return list
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

func fromHex(hex string) int {
	num, _ := strconv.ParseInt(hex, 16, 64)
	return int(num)
}
func toHex(num int) string {
	return fmt.Sprintf("%02x", num)
}

// return between 0 ~ 255 inclusive
func swingIncrement(base, step, gap int) int {
	MIN, MAX := 0, 255
	nextGap := gap
	prev := base
	upperLimit := false
	lowerLimit := false
	for i := 1; i <= step; i++ {
		if (i%2 == 1) {
			if tmp := prev + nextGap; upperLimit || tmp > MAX {
				upperLimit = true
				nextGap = gap
				if lowerLimit || prev-nextGap < MIN {
					break
				}
				prev = prev - nextGap
			} else {
				prev = prev + nextGap
			}
		} else {
			if tmp := prev - nextGap; lowerLimit || tmp < MIN {
				lowerLimit = true
				nextGap = gap
				if upperLimit || prev+nextGap > MAX {
					break
				}
				prev = prev + nextGap
			} else {
				prev = prev - nextGap
			}
		}
		if !upperLimit && !lowerLimit {
			nextGap += gap
		}
	}
	return prev
}
