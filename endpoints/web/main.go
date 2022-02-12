package main

import (
	"github.com/joho/godotenv"
	"github.com/yaroslavklimuk/crazy-lottery/server"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbFile := os.Getenv("DB_FILE")
	dataStorage, err := storage.GetStorage(dbFile)
	if err != nil {
		panic(err)
	}
	server.RegisterRoutes(dataStorage)
	server.RunServer()
}
