package main

import (
	"github.com/dannywolfmx/ReSender/server"
)

func main() {
	app := server.NewApp()
	app.Run()
}
