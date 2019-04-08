package server

import (
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/config"
	"os"
)

func Init(env string) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	initWeb(env)
}

func initWeb(env string) {
	router := NewRouter(env)

	port := os.Getenv("PORT") // provided by heroku's web dyno
	if port == "" {
		port = config.GetConfig().Get("port").(string)
	}

	_ = router.Run(":" + port)
}
