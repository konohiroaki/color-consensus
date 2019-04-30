package services

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type VoteService struct{}

func NewVoteService() VoteService {
	return VoteService{}
}

func (VoteService) Get(ctx *gin.Context, lang, name string, fields []string) []map[string]interface{} {
	return repositories.Vote(ctx).Get(lang, name, fields)
}

func (VoteService) Vote(ctx *gin.Context, lang, name string, newColors []string) {
	userID, _ := client.GetUserID(ctx)
	repositories.Vote(ctx).Add(userID, lang, name, newColors)
}

func (VoteService) RemoveByUser(ctx *gin.Context, userID string) {
	repositories.Vote(ctx).RemoveByUser(userID)
}
