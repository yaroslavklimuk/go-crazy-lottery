package server

import (
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"log"
	"net/http"
)

func RegisterRoutes(st storage.Storage) {
	handler := makeHttpHandler(st)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/get-reward", checkAuthMiddleware(handler.GetReward))
	http.HandleFunc("/submit-reward", checkAuthMiddleware(handler.SubmitReward))
	http.HandleFunc("/", checkAuthMiddleware(handler.Index))
}

func RunServer() {
	log.Fatal(http.ListenAndServe(":8081", nil))
}
