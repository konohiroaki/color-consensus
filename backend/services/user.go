package services

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (UserService) IsLoggedIn(ctx *gin.Context) bool {
	userID, err := client.GetUserID(ctx)

	return err == nil && repositories.User(ctx).IsPresent(userID)
}
