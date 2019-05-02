package client

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	getUserID = "github.com/konohiroaki/color-consensus/backend/client/getUserID"
	setUserID = "github.com/konohiroaki/color-consensus/backend/client/setUserID"
)

func GetUserIDFunc(ctx *gin.Context) (func() (string, error)) {
	return ctx.MustGet(getUserID).(func() (string, error))
}

func SetUserIDFunc(ctx *gin.Context) (func(string) error) {
	return ctx.MustGet(setUserID).(func(string) error)
}

func UserIDHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		getUserIDHandler,
		setUserIDHandler,
	}
}

func getUserIDHandler(ctx *gin.Context) {
	ctx.Set(getUserID, func() (string, error) {
		session := sessions.Default(ctx)
		userID := session.Get("userID")
		if userID != nil {
			return userID.(string), nil
		}
		return "", fmt.Errorf("user is not logged in")
	})
	ctx.Next()
}

func setUserIDHandler(ctx *gin.Context) {
	ctx.Set(setUserID, func(id string) error {
		session := sessions.Default(ctx)
		session.Set("userID", id)
		if err := session.Save(); err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	ctx.Next()
}
