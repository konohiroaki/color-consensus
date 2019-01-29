package main

import (
	"github.com/konohiroaki/color-consensus/backend/config"
	"github.com/konohiroaki/color-consensus/backend/server"
)

func main() {
	config.Init("development")

	server.Init()
}
