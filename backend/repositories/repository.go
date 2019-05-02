package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/config"
	"os"
	"strings"
)

const (
	voteKey     = "github.com/konohiroaki/color-consensus/backend/repositories/vote"
	userKey     = "github.com/konohiroaki/color-consensus/backend/repositories/user"
	languageKey = "github.com/konohiroaki/color-consensus/backend/repositories/language"
)

func Vote(ctx *gin.Context) VoteRepository {
	return ctx.MustGet(voteKey).(VoteRepository)
}

func User(ctx *gin.Context) UserRepository {
	return ctx.MustGet(userKey).(UserRepository)
}

func Language(ctx *gin.Context) LanguageRepository {
	return ctx.MustGet(languageKey).(LanguageRepository)
}

func Repositories(env string) []gin.HandlerFunc {
	uri, db := getDatabaseURIAndName()

	voteRepo := NewVoteRepository(uri, db, env)
	userRepo := NewUserRepository(uri, db, env)
	languageRepo := NewLanguageRepository()

	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			ctx.Set(voteKey, voteRepo)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(userKey, userRepo)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(languageKey, languageRepo)
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
