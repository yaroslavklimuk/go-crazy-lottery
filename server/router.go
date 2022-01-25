package server

import (
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"log"
	"net/http"
)

func RegisterRoutes(st storage.Storage) {
	http.HandleFunc("/register", makeRegisterHandler(st).ServeHTTP)
	http.HandleFunc("/login", makeLoginHandler(st).ServeHTTP)
	http.HandleFunc("/get-reward", checkAuthMiddleware(makeGetRewardHandler(st).ServeHTTP))
	http.HandleFunc("/submit-reward", checkAuthMiddleware(makeSubmitRewardHandler(st).ServeHTTP))
	http.HandleFunc("/", checkAuthMiddleware(makeIndexHandler(st).ServeHTTP))
}

func RunServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
