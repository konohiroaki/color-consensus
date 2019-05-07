package services

import (
	"fmt"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ColorService interface {
	GetAll() []map[string]interface{}
	Add(lang, name, code string, getUserID func() (string, error)) error
	GetNeighbors(code string, size int) ([]string, error)
	IsValidCodeFormat(input string) (bool, string)
}

type colorService struct {
	colorRepo repositories.ColorRepository
}

func NewColorService(colorRepo repositories.ColorRepository) ColorService {
	return colorService{colorRepo}
}

func (cs colorService) GetAll() []map[string]interface{} {
	return cs.colorRepo.GetAll([]string{"lang", "name", "code"})
}

// TODO: check lang existence in repo.
func (cs colorService) Add(lang, name, code string, getUserID func() (string, error)) error {
	userID, _ := getUserID()
	code = strings.ToLower(code)

	return cs.colorRepo.Add(lang, name, code, userID)
}

func (cs colorService) GetNeighbors(code string, size int) ([]string, error) {
	if 1 <= size && size <= 4096 {
		return cs.getNeighborColors(code, size), nil
	}
	return []string{}, fmt.Errorf("size should be between 1 and 4096")
}

func (colorService) IsValidCodeFormat(input string) (bool, string) {
	regex := regexp.MustCompile(`#[0-9a-fA-F]{6}`)
	return regex.MatchString(input), regex.String()
}

// TODO: I think the sort order shouldn't be measured only by diff scale but should also consider about the ratio between each RGB.
// For example of #808080, #707070 is farther than #806080, which would be opposite from how human feels.
// So currently it shows very bad result for gray-ish colors.
func (cs colorService) getNeighborColors(code string, size int) []string {
	r := cs.fromHex(code[0:2])
	g := cs.fromHex(code[2:4])
	b := cs.fromHex(code[4:6])
	type NeighborColor struct {
		Code string
		Diff int
	}
	candidates := []NeighborColor{{"#" + code, 0}}
	for i := 0; i < 256; i += 16 {
		for j := 0; j < 256; j += 16 {
			for k := 0; k < 256; k += 16 {
				diff := cs.abs(r-i) + cs.abs(g-j) + cs.abs(b-k)
				if diff != 0 {
					candidates = append(candidates, NeighborColor{
						"#" + cs.toHex(i) + cs.toHex(j) + cs.toHex(k),
						cs.abs(r-i) + cs.abs(g-j) + cs.abs(b-k),
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

func (colorService) fromHex(hex string) int {
	num, _ := strconv.ParseInt(hex, 16, 64)
	return int(num)
}
func (colorService) toHex(num int) string {
	return fmt.Sprintf("%02x", num)
}
func (colorService) abs(num int) int {
	if (num < 0) {
		return -num;
	}
	return num;
}
