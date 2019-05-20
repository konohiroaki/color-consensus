package services

import (
	"fmt"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"sort"
	"sync"
)

type ColorService interface {
	GetAll() []map[string]interface{}
	Add(category, name, code string, getUserID func() (string, error)) error
	GetNeighbors(code string, size int) ([]string, error)
}

type colorService struct {
	colorRepo         repositories.ColorRepository
	colorCategoryRepo repositories.ColorCategoryRepository
}

var (
	colorServiceInstance ColorService
	colorServiceOnce     sync.Once
)

func GetColorService(env string) ColorService {
	colorServiceOnce.Do(func() {
		colorServiceInstance = newColorService(repositories.GetColorRepository(env), repositories.GetColorCategoryRepository(env))
	})
	return colorServiceInstance
}

func newColorService(colorRepo repositories.ColorRepository, colorCategoryRepo repositories.ColorCategoryRepository) ColorService {
	return colorService{colorRepo, colorCategoryRepo}
}

func (cs colorService) GetAll() []map[string]interface{} {
	return cs.colorRepo.GetAll([]string{"category", "name", "code"})
}

func (cs colorService) Add(category, name, code string, getUserID func() (string, error)) error {
	userID, _ := getUserID()
	if !cs.colorCategoryRepo.IsPresent(category) {
		err := cs.colorCategoryRepo.Add(category, userID)
		if err != nil {
			return err
		}
	}

	code = util.shortToLowerLongHex(code)

	return cs.colorRepo.Add(category, name, code, userID)
}

func (cs colorService) GetNeighbors(code string, size int) ([]string, error) {
	if 1 <= size && size <= 4096 {
		return cs.getNeighborColors(code, size), nil
	}
	return []string{}, fmt.Errorf("size should be between 1 and 4096")
}

// TODO: I think the sort order shouldn't be measured only by diff scale but should also consider about the ratio between each RGB.
// For example of #808080, #707070 is farther than #806080, which would be opposite from how human feels.
// So currently it shows very bad result for gray-ish colors.
func (cs colorService) getNeighborColors(code string, size int) []string {
	r := util.fromHex(code[0:2])
	g := util.fromHex(code[2:4])
	b := util.fromHex(code[4:6])
	type NeighborColor struct {
		Code string
		Diff int
	}
	candidates := []NeighborColor{{"#" + code, 0}}
	for i := 0; i < 256; i += 16 {
		for j := 0; j < 256; j += 16 {
			for k := 0; k < 256; k += 16 {
				diff := util.abs(r-i) + util.abs(g-j) + util.abs(b-k)
				if diff != 0 {
					candidates = append(candidates, NeighborColor{
						"#" + util.toHex(i) + util.toHex(j) + util.toHex(k),
						util.abs(r-i) + util.abs(g-j) + util.abs(b-k),
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
