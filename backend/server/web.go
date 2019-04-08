package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/expvar"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/config"
	"github.com/konohiroaki/color-consensus/backend/controllers"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"os"
	"strings"
)

func NewRouter(env string) *gin.Engine {
	router := gin.Default()
	pprof.Register(router)

	router.GET("/debug/vars", expvar.Handler())
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("frontend/dist", false)))
	router.NoRoute(func(c *gin.Context) { c.File("frontend/dist/index.html") })
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("secret"))))
	router.Use(database(env)...)

	setUpEndpoints(router)

	return router
}

func database(env string) []gin.HandlerFunc {
	uri, db := getDatabaseURIAndName()
	colorRepository := repositories.NewColorRepository(uri, db, env)
	voteRepository := repositories.NewVoteRepository(uri, db, env)
	userRepository := repositories.NewUserRepository(uri, db, env)

	return []gin.HandlerFunc{
		func(c *gin.Context) {
			c.Set("colorRepository", colorRepository)
			c.Next()
		},
		func(c *gin.Context) {
			c.Set("voteRepository", voteRepository)
			c.Next()
		},
		func(c *gin.Context) {
			c.Set("userRepository", userRepository)
			c.Next()
		},
	}
}

func getDatabaseURIAndName() (string, string) {
	uri := os.Getenv("MONGODB_URI") // provided by mLab add-on
	db := uri[strings.LastIndex(uri, "/")+1:]
	if uri == "" {
		uri = config.GetConfig().Get("mongo.url").(string)
		db = "cc"
	}
	return uri, db
}

func setUpEndpoints(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1api := api.Group("/v1")
		{
			color := new(controllers.ColorController)
			v1api.GET("/colors", color.GetAll)
			v1api.POST("/colors", color.Add)
			// TODO: change path to /colors/:code/neighbors
			//v1api.GET("/colors/candidates/:code", color.GetNeighbors)
			v1api.GET("/colors/:code/neighbors", color.GetNeighbors)

			vote := new(controllers.VoteController)
			v1api.POST("/votes", vote.Vote)
			v1api.GET("/votes", vote.GetVotes)
			v1api.GET("/votes-stats", vote.GetStats)
			v1api.DELETE("/votes", vote.DeleteVotesForUser)

			userController := new(controllers.UserController)
			v1api.GET("/users/presence", userController.GetUserIDFromCookie)
			v1api.POST("/users/presence", userController.SetCookieIfUserExist)
			v1api.POST("/users", userController.AddUserAndSetCookie)
		}
	}
}
