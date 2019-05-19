package client

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type Client interface {
	GetUserIDFunc(ctx *gin.Context) (func() (string, error))
	SetUserIDFunc(ctx *gin.Context) (func(string) error)
}

type client struct{}

var (
	clientInstance Client
	clientOnce     sync.Once
)

func GetClient() Client {
	clientOnce.Do(func() {
		client := client{}
		clientInstance = client
	})
	return clientInstance
}

func (client) GetUserIDFunc(ctx *gin.Context) (func() (string, error)) {
	return func() (string, error) {
		session := sessions.Default(ctx)
		userID := session.Get("userID")
		if userID != nil {
			return userID.(string), nil
		}
		return "", fmt.Errorf("user is not logged in")
	}
}

func (client) SetUserIDFunc(ctx *gin.Context) (func(string) error) {
	return func(id string) error {
		session := sessions.Default(ctx)
		session.Set("userID", id)
		if err := session.Save(); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}
