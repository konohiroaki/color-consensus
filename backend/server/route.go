package server

import (
	"github.com/gin-contrib/cors"
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
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", false)))
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	api := router.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			consensus := new(controllers.ConsensusController)
			v1api.GET("/colors/keys", consensus.GetAllConsensusKey)
			v1api.GET("/colors/detail", consensus.GetAllConsensus)
			v1api.GET("/colors/keys/:lang", consensus.GetAllConsensusKeyForLang)
			v1api.GET("/colors/detail/:lang", consensus.GetAllConsensusForLang)
			v1api.GET("/colors/detail/:lang/:color", consensus.GetConsensus)
			v1api.GET("/colors/candidates/:code", consensus.GetCandidateList)

			vote := new(controllers.VoteController)
			v1api.POST("/votes/:lang/:color", vote.Vote)
			v1api.GET("/votes/:lang/:color/raw", vote.GetVotes)

			user := new(controllers.UserController)
			v1api.GET("/users/presence", user.GetPresence)
			v1api.POST("/users", user.RegisterUser)
			v1api.GET("/admin/users", user.GetUserList)
		}
	}
	return router
}
