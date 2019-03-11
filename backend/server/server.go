package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/konohiroaki/color-consensus/backend/config"
	"github.com/konohiroaki/color-consensus/backend/domains/consensus"
	"github.com/konohiroaki/color-consensus/backend/domains/user"
	"github.com/konohiroaki/color-consensus/backend/domains/vote"
	"os"
	"strings"
)

func Init(env string) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	initRepo(env)
	initWeb()
}

func initRepo(env string) {
	uri := os.Getenv("MONGODB_URI") // provided by mLab add-on
	db := uri[strings.LastIndex(uri, "/")+1:]
	if uri == "" {
		uri = config.GetConfig().Get("mongo.url").(string)
		db = "cc"
	}

	user.InitRepo(uri, db)
	vote.InitRepo(uri, db)
	consensus.InitRepo(uri, db)

	if env == "development" {
		fmt.Println("detected development mode. inserting sample data.")
		insertSampleData()
	}
}

func initWeb() {
	router := NewRouter()

	port := os.Getenv("PORT") // provided by heroku's web dyno
	if port == "" {
		port = config.GetConfig().Get("port").(string)
	}

	_ = router.Run(":" + port)
}

func insertSampleData() {
	user.InsertSampleData()
	vote.InsertSampleData()
	consensus.InsertSampleData()
}
