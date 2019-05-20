package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"sync"
)

type VoteService interface {
	Get(category, name string, fields []string) []map[string]interface{}
	Vote(category, name string, newColors []string, getUserID func() (string, error))
	RemoveByUser(userID string)
}

type voteService struct {
	voteRepo repositories.VoteRepository
}

var (
	voteServiceInstance VoteService
	voteServiceOnce     sync.Once
)

func GetVoteService(env string) VoteService {
	voteServiceOnce.Do(func() {
		voteServiceInstance = newVoteService(repositories.GetVoteRepository(env))
	})
	return voteServiceInstance
}

func newVoteService(voteRepo repositories.VoteRepository) VoteService {
	return voteService{voteRepo}
}

func (vs voteService) Get(category, name string, fields []string) []map[string]interface{} {
	return vs.voteRepo.Get(category, name, fields)
}

func (vs voteService) Vote(category, name string, newColors []string, getUserID func() (string, error)) {
	userID, _ := getUserID() // ensured there is no error

	for k := range newColors {
		newColors[k] = util.shortToLowerLongHex(newColors[k])
	}

	vs.voteRepo.Add(category, name, newColors, userID)
}

func (vs voteService) RemoveByUser(userID string) {
	vs.voteRepo.RemoveByUser(userID)
}
