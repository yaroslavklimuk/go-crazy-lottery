package server

import (
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"html/template"
	"net/http"
)

type (
	RandomRewardResponse struct {
		TaskId string `json:"task_id"`
	}

	baseHttpHandler struct {
		storage storage.Storage
	}
	indexRequestHandler struct {
		baseHttpHandler
	}
	registerRequestHandler struct {
		baseHttpHandler
	}
	loginRequestHandler struct {
		baseHttpHandler
	}
	getRewardRequestHandler struct {
		baseHttpHandler
	}
	submitRewardRequestHandler struct {
		baseHttpHandler
	}
)

func (h *indexRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		renderTemplate("./templates/index.html", writer, nil)
	default:
		http.Error(writer, "Sorry, only GET requests are supported.", 405)
		return
	}
}

func (h *registerRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

func (h *loginRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

func (h *getRewardRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

func (h *submitRewardRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/submit-reward" {
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

func makeIndexHandler(st storage.Storage) http.Handler {
	return &indexRequestHandler{struct{ storage storage.Storage }{storage: st}}
}

func makeRegisterHandler(st storage.Storage) http.Handler {
	return &registerRequestHandler{struct{ storage storage.Storage }{storage: st}}
}

func makeLoginHandler(st storage.Storage) http.Handler {
	return &loginRequestHandler{struct{ storage storage.Storage }{storage: st}}
}

func makeGetRewardHandler(st storage.Storage) http.Handler {
	return &getRewardRequestHandler{struct{ storage storage.Storage }{storage: st}}
}

func makeSubmitRewardHandler(st storage.Storage) http.Handler {
	return &submitRewardRequestHandler{struct{ storage storage.Storage }{storage: st}}
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
