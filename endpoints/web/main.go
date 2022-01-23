package main

import (
	"github.com/yaroslavklimuk/crazy-lottery/server"
)

func main() {
	server.RegisterRoutes()
	server.RunServer()
}
