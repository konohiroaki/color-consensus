package services

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

const (
	colorKey    = "github.com/konohiroaki/color-consensus/backend/services/color"
	voteKey     = "github.com/konohiroaki/color-consensus/backend/services/vote"
	userKey     = "github.com/konohiroaki/color-consensus/backend/services/user"
	languageKey = "github.com/konohiroaki/color-consensus/backend/services/language"
)

func Color(ctx *gin.Context) ColorService {
	return ctx.MustGet(colorKey).(ColorService)
}

func Vote(ctx *gin.Context) VoteService {
	return ctx.MustGet(voteKey).(VoteService)
}

func User(ctx *gin.Context) UserService {
	return ctx.MustGet(userKey).(UserService)
}

func Language(ctx *gin.Context) LanguageService {
	return ctx.MustGet(languageKey).(LanguageService)
}

func Services(env string) []gin.HandlerFunc {
	colorRepo := repositories.NewColorRepository(env)

	colorService := NewColorService(colorRepo)
	voteService := NewVoteService()
	userService := NewUserService()
	languageService := NewLanguageService()

	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			ctx.Set(colorKey, colorService)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(voteKey, voteService)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(userKey, userService)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(languageKey, languageService)
			ctx.Next()
		},
	}
}
