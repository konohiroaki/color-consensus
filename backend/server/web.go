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

	setUpEndpoints(router, env)

	return router
}

func setUpEndpoints(router *gin.Engine, env string) {
	color, vote, user, language := getControllers(env)

	api := router.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			v1api.POST("/colors", color.Add)
			v1api.GET("/colors", color.GetAll)
			v1api.GET("/colors/:code/neighbors", color.GetNeighbors)

			v1api.POST("/votes", vote.Vote)
			v1api.GET("/votes", vote.Get)

			v1api.POST("/users", user.SingUpAndLogin)
			v1api.POST("/login", user.Login)
			v1api.GET("/users/presence", user.GetIDIfLoggedIn)

			v1api.GET("/languages", language.GetAll)
		}
	}
}

func getControllers(env string) (color controllers.ColorController, vote controllers.VoteController, user controllers.UserController, language controllers.LanguageController) {
	colorRepo := repositories.NewColorRepository(env)
	voteRepo := repositories.NewVoteRepository(env)
	userRepo := repositories.NewUserRepository(env)
	langRepo := repositories.NewLanguageRepository()

	colorService := services.NewColorService(colorRepo)
	voteService := services.NewVoteService(voteRepo)
	userService := services.NewUserService(userRepo)
	langService := services.NewLanguageService(langRepo)

	color = controllers.NewColorController(colorService, userService)
	vote = controllers.NewVoteController(voteService, userService)
	user = controllers.NewUserController(userService)
	language = controllers.NewLanguageController(langService)

	return
}
