package server

import (
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/register", handleRegisterRequest)
	http.HandleFunc("/login", handleLoginRequest)
	http.HandleFunc("/get-reward", checkAuthMiddleware(handleRewardRequest))
	http.HandleFunc("/submit-reward", checkAuthMiddleware(handleSubmitRewardRequest))
	http.HandleFunc("/", handleIndexRequest)
}

func RunServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
