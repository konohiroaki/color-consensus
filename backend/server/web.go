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

func NewRouter(env string) *gin.Engine {
	router := gin.Default()

	if env == "development" {
		pprof.Register(router)
		router.GET("/debug/vars", expvar.Handler())
		router.Use(cors.Default())
	}

	router.Use(static.Serve("/", static.LocalFile("frontend/dist", false)))
	router.NoRoute(func(c *gin.Context) { c.File("frontend/dist/index.html") })
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))

	setUpEndpoints(router, env)

	return router
}

func setUpEndpoints(router *gin.Engine, env string) {
	api := router.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			color := controllers.GetColorController(env)
			v1api.POST("/colors", color.Add)
			v1api.GET("/colors", color.GetAll)
			v1api.GET("/colors/:code/neighbors", color.GetNeighbors)

			vote := controllers.GetVoteController(env)
			v1api.POST("/votes", vote.Vote)
			v1api.GET("/votes", vote.Get)

			user := controllers.GetUserController(env)
			v1api.POST("/users/signup", user.SignUpAndLogin)
			v1api.POST("/users/login", user.Login)
			v1api.GET("/users", user.GetIDIfLoggedIn)

			language := controllers.GetLanguageController()
			v1api.GET("/languages", language.GetAll)
			colorCategory := controllers.GetColorCategoryController(env)
			v1api.GET("/color-categories", colorCategory.GetAll)
			nationality := controllers.GetNationalityController()
			v1api.GET("/nationalities", nationality.GetAll)
			gender := controllers.GetGenderController()
			v1api.GET("/genders", gender.GetAll)
		}
	}
}
