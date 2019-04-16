package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"net/http"
	"sort"
	"strconv"
)

type ColorController struct{}

func (ColorController) GetAll(ctx *gin.Context) {
	repository := ctx.Keys["colorRepository"].(repositories.ColorRepository)
	ctx.JSON(200, repository.GetAll([]string{"lang", "name", "code"}))
}

func (ColorController) Add(ctx *gin.Context) {
	repository := ctx.Keys["colorRepository"].(repositories.ColorRepository)
	user := ctx.Keys["userRepository"].(repositories.UserRepository)
	session := sessions.Default(ctx)
	userID := session.Get("userID")
	if userID == nil || !user.IsPresent(userID.(string)) {
		fmt.Println(userID)
		ctx.Status(http.StatusForbidden)
		return
	}
	type request struct {
		Lang string `json:"lang"`
		Name string `json:"name"`
		Code string `json:"code"`
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	repository.Add(req.Lang, req.Name, req.Code, userID.(string))
	ctx.Status(http.StatusCreated);
}

func (ColorController) GetNeighbors(ctx *gin.Context) {
	code := ctx.Param("code")
	if size, err := strconv.Atoi(ctx.Query("size")); err == nil {
		neighbors := getNeighborColors(code, size)
		ctx.JSON(200, neighbors)
	} else {
		ctx.AbortWithStatus(400)
	}
}

// TODO: I think the sort order shouldn't be measured only by diff scale but should also consider about the ratio between each RGB.
// For example of #808080, #707070 is farther than #806080, which would be opposite from how human feels.
// So currently it shows very bad result for gray-ish colors.
func getNeighborColors(code string, size int) []string {
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
				diff := abs(r-i) + abs(g-j) + abs(b-k)
				if diff != 0 {
					list = append(list, Candidate{
						"#" + toHex(i) + toHex(j) + toHex(k),
						abs(r-i) + abs(g-j) + abs(b-k),
					})
				}
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
