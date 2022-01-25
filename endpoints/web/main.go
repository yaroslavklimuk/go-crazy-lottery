package main

import (
	"github.com/joho/godotenv"
	"github.com/yaroslavklimuk/crazy-lottery/server"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}
	dbFile := env["DB_FILE"]
	dataStorage, err := storage.GetStorage(dbFile)
	if err != nil {
		panic(err)
	}
	server.RegisterRoutes(dataStorage)
	server.RunServer()
}
