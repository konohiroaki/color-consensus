package server

import (
	"fmt"
	"github.com/konohiroaki/color-consensus/backend/config"
	"github.com/konohiroaki/color-consensus/backend/domains/consensus"
	"github.com/konohiroaki/color-consensus/backend/domains/user"
	"github.com/konohiroaki/color-consensus/backend/domains/vote"
	"os"
)

func Init(env string) {
	initRepo(env)

	router := NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = config.GetConfig().Get("port").(string)
	}
	_ = router.Run(":" + port)
}

func initRepo(env string) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = config.GetConfig().Get("mongo.url").(string)
	}
	db := "cc"

	user.InitRepo(uri, db)
	vote.InitRepo(uri, db)
	consensus.InitRepo(uri, db)

	if env == "development" {
		fmt.Println("detected development mode. inserting sample data.")
		insertSampleData()
	}
}

func insertSampleData() {
	user.InsertSampleData()
	vote.InsertSampleData()
	consensus.InsertSampleData()
}
