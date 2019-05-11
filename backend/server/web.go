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

	setUpEndpoints(router, env)

	return router
}

func setUpEndpoints(router *gin.Engine, env string) {
	color, vote, user, language, nationality, gender := getControllers(env)

	api := router.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			v1api.POST("/colors", color.Add)
			v1api.GET("/colors", color.GetAll)
			v1api.GET("/colors/:code/neighbors", color.GetNeighbors)

			v1api.POST("/votes", vote.Vote)
			v1api.GET("/votes", vote.Get)

			v1api.POST("/users/signup", user.SignUpAndLogin)
			v1api.POST("/users/login", user.Login)
			v1api.GET("/users", user.GetIDIfLoggedIn)

			v1api.GET("/languages", language.GetAll)
			v1api.GET("/nationalities", nationality.GetAll)
			v1api.GET("/genders", gender.GetAll)
		}
	}
}

func getControllers(env string) (color controllers.ColorController, vote controllers.VoteController,
		user controllers.UserController, language controllers.LanguageController,
		nationality controllers.NationalityController, gender controllers.GenderController) {
	colorRepo := repositories.NewColorRepository(env)
	voteRepo := repositories.NewVoteRepository(env)
	userRepo := repositories.NewUserRepository(env)
	langRepo := repositories.NewLanguageRepository()
	nationRepo := repositories.NewNationalityRepository()
	genderRepo := repositories.NewGenderRepository()

	clientHandler := client.NewClient()
	colorService := services.NewColorService(colorRepo)
	voteService := services.NewVoteService(voteRepo)
	userService := services.NewUserService(userRepo, nationRepo, genderRepo)
	langService := services.NewLanguageService(langRepo)
	nationService := services.NewNationalityService(nationRepo)
	genderService := services.NewGenderService(genderRepo)

	color = controllers.NewColorController(colorService, userService, clientHandler)
	vote = controllers.NewVoteController(voteService, userService, clientHandler)
	user = controllers.NewUserController(userService, clientHandler)
	language = controllers.NewLanguageController(langService)
	nationality = controllers.NewNationalityController(nationService)
	gender = controllers.NewGenderController(genderService)

	return
}
