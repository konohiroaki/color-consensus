package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	repo "github.com/konohiroaki/color-consensus/backend/repositories"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type ColorController struct{}

func (ColorController) GetAll(ctx *gin.Context) {
	colorRepo := repo.Color(ctx)
	ctx.JSON(http.StatusOK, colorRepo.GetAll([]string{"lang", "name", "code"}))
}

func (ColorController) Add(ctx *gin.Context) {
	colorRepo := repo.Color(ctx)
	userRepo := repo.User(ctx)
	userID, err := client.GetUserID(ctx)
	if err != nil || !userRepo.IsPresent(userID) {
		log.Println(err, userID)
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
	colorRepo.Add(req.Lang, req.Name, req.Code, userID)
	ctx.Status(http.StatusCreated);
}

func (ColorController) GetNeighbors(ctx *gin.Context) {
	code := ctx.Param("code")
	if size, err := strconv.Atoi(ctx.Query("size")); err == nil {
		neighbors := getNeighborColors(code, size)
		ctx.JSON(http.StatusOK, neighbors)
	} else {
		ctx.Status(http.StatusBadRequest)
	}
}

// TODO: I think the sort order shouldn't be measured only by diff scale but should also consider about the ratio between each RGB.
// For example of #808080, #707070 is farther than #806080, which would be opposite from how human feels.
// So currently it shows very bad result for gray-ish colors.
func getNeighborColors(code string, size int) []string {
	r := fromHex(code[0:2])
	g := fromHex(code[2:4])
	b := fromHex(code[4:6])
	type NeighborColor struct {
		Code string
		Diff int
	}
	candidates := []NeighborColor{{"#" + code, 0}}
	for i := 0; i < 256; i += 16 {
		for j := 0; j < 256; j += 16 {
			for k := 0; k < 256; k += 16 {
				diff := abs(r-i) + abs(g-j) + abs(b-k)
				if diff != 0 {
					candidates = append(candidates, NeighborColor{
						"#" + toHex(i) + toHex(j) + toHex(k),
						abs(r-i) + abs(g-j) + abs(b-k),
					})
				}
			}
		}
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i].Diff < candidates[j].Diff })
	result := []string{}
	for _, candidate := range candidates {
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

func errorResponse(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"message": message,
		},
	}
}
