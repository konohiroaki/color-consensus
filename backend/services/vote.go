package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type VoteService interface {
	Get(lang, name string, fields []string) []map[string]interface{}
	Vote(lang, name string, newColors []string, getUserID func() (string, error))
	RemoveByUser(userID string)
}

type voteService struct {
	voteRepo repositories.VoteRepository
}

func NewVoteService(voteRepo repositories.VoteRepository) VoteService {
	return voteService{voteRepo}
}

func (vs voteService) Get(lang, name string, fields []string) []map[string]interface{} {
	return vs.voteRepo.Get(lang, name, fields)
}

func (vs voteService) Vote(lang, name string, newColors []string, getUserID func() (string, error)) {
	userID, _ := getUserID()
	vs.voteRepo.Add(userID, lang, name, newColors)
}

func (vs voteService) RemoveByUser(userID string) {
	vs.voteRepo.RemoveByUser(userID)
}
