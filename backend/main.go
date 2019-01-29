package main

import (
	"flag"
	"github.com/konohiroaki/color-consensus/backend/config"
	"github.com/konohiroaki/color-consensus/backend/server"
)

func main() {
	env := flag.String("env","production", "specify environment")
	flag.Parse()
	config.Init(*env)

	server.Init()
}
