package services

import "github.com/gin-gonic/gin"

const (
	colorKey    = "github.com/konohiroaki/color-consensus/backend/service/color"
	voteKey     = "github.com/konohiroaki/color-consensus/backend/service/vote"
	userKey     = "github.com/konohiroaki/color-consensus/backend/service/user"
	languageKey = "github.com/konohiroaki/color-consensus/backend/service/language"
)

func Color(ctx *gin.Context) ColorService {
	return ctx.MustGet(colorKey).(ColorService)
}

func User(ctx *gin.Context) UserService {
	return ctx.MustGet(userKey).(UserService)
}

func Services() []gin.HandlerFunc {
	colorService := NewColorService()
	//voteService := NewVoteService()
	userService := NewUserService()
	//languageService := NewLanguageService()

	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			ctx.Set(colorKey, colorService)
			ctx.Next()
		},
		func(ctx *gin.Context) {
			ctx.Set(userKey, userService)
			ctx.Next()
		},
	}
}
