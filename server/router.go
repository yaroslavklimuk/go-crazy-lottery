package server

import (
	"github.com/yaroslavklimuk/crazy-lottery/http_handler"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"log"
	"net/http"
)

func RegisterRoutes(st storage.Storage) {
	handler := http_handler.MakeHttpHandler(st)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/get-reward", handler.GetReward)
	http.HandleFunc("/submit-reward", handler.SubmitReward)
	http.HandleFunc("/", handler.Index)
}

func RunServer() {
	log.Fatal(http.ListenAndServe(":8081", nil))
}
