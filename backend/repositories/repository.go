package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/config"
	"os"
	"strings"
)

const (
	colorKey = "github.com/konohiroaki/color-consensus/backend/repositories/color"
	voteKey  = "github.com/konohiroaki/color-consensus/backend/repositories/vote"
	userKey  = "github.com/konohiroaki/color-consensus/backend/repositories/user"
)

func Color(ctx *gin.Context) ColorRepository {
	return ctx.MustGet(colorKey).(ColorRepository)
}

func Vote(ctx *gin.Context) VoteRepository {
	return ctx.MustGet(voteKey).(VoteRepository)
}

func User(ctx *gin.Context) UserRepository {
	return ctx.MustGet(userKey).(UserRepository)
}

func Repositories(env string) []gin.HandlerFunc {
	uri, db := getDatabaseURIAndName()

	colorRepo := NewColorRepository(uri, db, env)
	voteRepo := NewVoteRepository(uri, db, env)
	userRepo := NewUserRepository(uri, db, env)

	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			ctx.Set(colorKey, colorRepo)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(voteKey, voteRepo)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(userKey, userRepo)
			ctx.Next()
		},
	}
}

func getDatabaseURIAndName() (string, string) {
	uri := os.Getenv("MONGODB_URI") // provided by mLab add-on
	db := uri[strings.LastIndex(uri, "/")+1:]
	if uri == "" {
		uri = config.GetConfig().Get("mongo.url").(string)
		db = "cc"
	}
	return uri, db
}
