package main

import (
	"github.com/konohiroaki/color-consensus/backend/route"
	"os"
)

func main() {
	Init()
}

func Init() {
	router := route.Route()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router.Run(":" + port)
}
