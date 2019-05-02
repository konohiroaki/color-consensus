package services

import (
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type VoteService struct {
	voteRepo repositories.VoteRepository
}

func NewVoteService(voteRepo repositories.VoteRepository) VoteService {
	return VoteService{voteRepo}
}

func (vs VoteService) Get(lang, name string, fields []string) []map[string]interface{} {
	return vs.voteRepo.Get(lang, name, fields)
}

func (vs VoteService) Vote(lang, name string, newColors []string, getUserID func() (string, error)) {
	userID, _ := getUserID()
	vs.voteRepo.Add(userID, lang, name, newColors)
}

func (vs VoteService) RemoveByUser(userID string) {
	vs.voteRepo.RemoveByUser(userID)
}
