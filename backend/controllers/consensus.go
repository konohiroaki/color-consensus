package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/models"
	"net/http"
	"sort"
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
		list = append(list, ResponseElement{e.Language, e.Color, e.Code})
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
			list = append(list, ResponseElement{e.Language, e.Color, e.Code})
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
	if size, err := strconv.Atoi(c.Query("size")); err == nil {
		candidates := generateCandidateList(code, size)
		c.JSON(200, candidates)
	} else {
		c.AbortWithStatus(400)
	}
}

func (ConsensusController) AddColor(c *gin.Context) {
	var colorConsensus models.ColorConsensus
	// TODO: error handling.
	c.BindJSON(&colorConsensus)
	colorConsensus.Colors = map[string]int{}
	colorConsensus.Vote = 0
	models.Consensus = append(models.Consensus, &colorConsensus);
	c.Status(http.StatusCreated);
}

// FIXME: the result doesn't look nice.
func generateCandidateList(code string, size int) []string {
	r := fromHex(code[0:2])
	g := fromHex(code[2:4])
	b := fromHex(code[4:6])
	type Candidate struct {
		Code string
		Diff int
	}
	list := []Candidate{{"#" + code, 0}}
	for i := 0; i < 256; i += 16 {
		for j := 0; j < 256; j += 16 {
			for k := 0; k < 256; k += 16 {
				list = append(list, Candidate{
					"#" + toHex(i) + toHex(j) + toHex(k),
					abs(r-i) + abs(g-j) + abs(b-k),
				})
			}
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Diff < list[j].Diff })
	result := []string{}
	for _, candidate := range list {
		result = append(result, candidate.Code)
		if (len(result) == size) {
			break;
		}
	}
	return result
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
func abs(num int) int {
	if (num < 0) {
		return -num;
	}
	return num;
}
