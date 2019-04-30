package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

func (UserService) IsLoggedIn(ctx *gin.Context) bool {
	userID, err := client.GetUserID(ctx)

	return err == nil && repositories.User(ctx).IsPresent(userID)
}

func (UserService) GetID(ctx *gin.Context) (string, error) {
	userID, err := client.GetUserID(ctx)

	if err != nil || !repositories.User(ctx).IsPresent(userID) {
		return "", fmt.Errorf("user is not logged in")
	}

	return userID, nil
}

func (UserService) SingUpAndLogin(ctx *gin.Context, nationality, gender string, birth int) (string, bool) {
	id := repositories.User(ctx).Add(nationality, gender, birth)

	cookieErr := client.SetUserID(ctx, id)
	if cookieErr != nil {
		// ignore remove error because rare and it doesn't harm.
		_ = repositories.User(ctx).Remove(id)
		return "", false
	}

	return id, true
}

func (UserService) TryLogin(ctx *gin.Context, userID string) bool {
	if repositories.User(ctx).IsPresent(userID) {
		err := client.SetUserID(ctx, userID)
		return err == nil
	}
	return false
}
