package server

import (
	"fmt"
	"github.com/konohiroaki/color-consensus/backend/config"
	"os"
)

func Init() {
	tryConfig()

	router := NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = config.GetConfig().Get("port").(string)
	}
	router.Run(":" + port)
}

func tryConfig() {
	fmt.Println(config.GetConfig().Get("test"))
}
