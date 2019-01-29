package server

import (
	"github.com/gin-contrib/expvar"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	pprof.Register(router)

	router.GET("/debug/vars", expvar.Handler())
	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", false)))
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	api := router.Group("/api")
	{
		color := new(controllers.ColorController)
		api.GET("/colors", color.GetAllConsensus)
		api.GET("/colors/:lang", color.GetAllConsensusForLang)
		api.GET("/colors/:lang/:color", color.GetConsensus)

		vote := new(controllers.VoteController)
		api.POST("/colors/:lang/:color/:user", vote.Vote)
		api.GET("/colors/:lang/:color/raw", vote.GetVotes)

		user := new(controllers.UserController)
		api.GET("/users/presence", user.GetPresence)
		api.POST("/users", user.RegisterUser)
		api.GET("/admin/users", user.GetUserList)
	}
	return router
}
