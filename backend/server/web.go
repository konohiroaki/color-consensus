package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/expvar"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/client"
	"github.com/konohiroaki/color-consensus/backend/controllers"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"github.com/konohiroaki/color-consensus/backend/services"
)

func NewRouter(env string) *gin.Engine {
	router := gin.Default()
	pprof.Register(router)

	router.GET("/debug/vars", expvar.Handler())
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("frontend/dist", false)))
	router.NoRoute(func(c *gin.Context) { c.File("frontend/dist/index.html") })
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))
	router.Use(client.UserIDHandlers()...)
	router.Use(services.Services(env)...)
	router.Use(repositories.Repositories(env)...)

	setUpEndpoints(router)

	return router
}

func setUpEndpoints(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			color := controllers.NewColorController()
			v1api.POST("/colors", color.Add)
			v1api.GET("/colors", color.GetAll)
			v1api.GET("/colors/:code/neighbors", color.GetNeighbors)

			vote := controllers.NewVoteController()
			v1api.POST("/votes", vote.Vote)
			v1api.GET("/votes", vote.Get)

			userController := controllers.NewUserController()
			v1api.POST("/users", userController.SingUpAndLogin)
			v1api.POST("/login", userController.Login)
			v1api.GET("/users/presence", userController.GetIDIfLoggedIn)

			langController := controllers.NewLanguageController()
			v1api.GET("/languages", langController.GetAll)
		}
	}
}

func getControllers() {
	color := controllers.NewColorController()
	vote := controllers.NewVoteController()
	user := controllers.NewUserController()
	language := controllers.NewLanguageController()
}
