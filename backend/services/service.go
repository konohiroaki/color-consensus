package services

import "github.com/gin-gonic/gin"

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

func Services() []gin.HandlerFunc {
	colorService := NewColorService()
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
