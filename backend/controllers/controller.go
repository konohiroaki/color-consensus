package controllers

import "github.com/gin-gonic/gin"

const (
	colorKey    = "github.com/konohiroaki/color-consensus/backend/controllers/color"
	voteKey     = "github.com/konohiroaki/color-consensus/backend/controllers/vote"
	userKey     = "github.com/konohiroaki/color-consensus/backend/controllers/user"
	languageKey = "github.com/konohiroaki/color-consensus/backend/controllers/language"
)

func Color(ctx *gin.Context) ColorController {
	return ctx.MustGet(colorKey).(ColorController)
}

func Controllers() []gin.HandlerFunc {
	colorController := NewColorController()

	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			ctx.Set(colorKey, colorController)
			ctx.Next()
		},
	}
}
