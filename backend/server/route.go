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
	"github.com/konohiroaki/color-consensus/backend/domains/user"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	pprof.Register(router)

	router.GET("/debug/vars", expvar.Handler())
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("frontend/dist", false)))
	router.NoRoute(func(c *gin.Context) { c.File("frontend/dist/index.html") })
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	api := router.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			consensus := new(controllers.ConsensusController)
			v1api.GET("/colors/keys", consensus.GetAllConsensusKey)
			v1api.GET("/colors/detail", consensus.GetAllConsensus)
			v1api.GET("/colors/detail/:lang/:color", consensus.GetConsensus)
			v1api.GET("/colors/candidates/:code", consensus.GetCandidateList)
			v1api.POST("/colors", consensus.AddColor)

			vote := new(controllers.VoteController)
			v1api.POST("/votes", vote.Vote)
			v1api.GET("/votes/:lang/:color/raw", vote.GetVotes)

			userController := new(controllers.UserController)
			v1api.GET("/users/presence", userController.GetUserIDFromCookie)
			v1api.POST("/users/presence", userController.SetCookieIfUserExist)
			v1api.POST("/users", userController.AddUserAndSetCookie)
			v1api.GET("/admin/users", func(c *gin.Context) { c.JSON(200, user.GetList()) })
		}
	}
	return router
}
