package server

import (
	"net/http"
)

type (
	RandomRewardResponse struct {
		TaskId string `json:"task_id"`
	}
)

func handleIndexRequest(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		writer.Header().Set("Content-Type", "text/html")
	default:
		http.Error(writer, "Sorry, only GET requests are supported.", 405)
		return
	}
}

func handleRegisterRequest(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/register" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "POST":
		writer.Header().Set("Content-Type", "application/json")
	default:
		http.Error(writer, "Sorry, only POST requests are supported.", 405)
		return
	}
}

func handleLoginRequest(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/login" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "POST":
		writer.Header().Set("Content-Type", "application/json")
	default:
		http.Error(writer, "Sorry, only POST requests are supported.", 405)
		return
	}
}

func handleRewardRequest(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/get-reward" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "POST":
		writer.Header().Set("Content-Type", "application/json")
	default:
		http.Error(writer, "Sorry, only POST requests are supported.", 405)
		return
	}
}

func handleSubmitRewardRequest(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/get-reward" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "POST":
		writer.Header().Set("Content-Type", "application/json")
	default:
		http.Error(writer, "Sorry, only POST requests are supported.", 405)
		return
	}
}
