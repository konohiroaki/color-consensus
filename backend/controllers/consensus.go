package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/domains/consensus"
	"net/http"
	"sort"
	"strconv"
)

type ConsensusController struct{}

func (ConsensusController) GetAllConsensusKey(c *gin.Context) {
	fmt.Println(consensus.GetKeys())
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
	var colorConsensus consensus.ColorConsensus
	_ = c.BindJSON(&colorConsensus)
	consensus.Add(colorConsensus)

	c.Status(http.StatusCreated);
}

// TODO: I think the sort order shouldn't be measured only by diff scale but should also consider about the ratio between each RGB.
// It shows very bad result for gray.
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
