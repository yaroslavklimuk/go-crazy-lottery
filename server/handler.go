package server

import (
	"html/template"
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
		renderTemplate("./templates/index.html", writer, nil)
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
	case "GET":
		renderTemplate("./templates/register.html", writer, nil)
	case "POST":
		writer.Header().Set("Content-Type", "text/html")
	default:
		http.Error(writer, "Sorry, only GET or POST requests are supported.", 405)
		return
	}
}

func handleLoginRequest(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/login" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "GET":
		renderTemplate("./templates/login.html", writer, nil)
	case "POST":
		writer.Header().Set("Content-Type", "application/json")
	default:
		http.Error(writer, "Sorry, only GET or POST requests are supported.", 405)
		return
	}
}

func renderTemplate(templName string, writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "text/html")
	ts, err := template.ParseFiles(templName)
	if err != nil {
		http.Error(writer, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Internal Server Error", 500)
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
