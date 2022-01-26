package server

import (
	"github.com/google/uuid"
	"github.com/yaroslavklimuk/crazy-lottery/dto"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"html/template"
	"net/http"
	"time"
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
		form := request.PostForm
		user, err := h.storage.GetUserByName(form.Get("login"))
		if err != nil {
			returnAjaxError(writer, err.Error(), 500)
			return
		}
		if user != nil {
			returnAjaxError(writer, "User already exists", 400)
			return
		}
		newUser := dto.NewUser(
			0,
			form.Get("login"),
			form.Get("banc_acc"),
			form.Get("address"),
			0,
			form.Get("passwd"),
		)
		userId, err := h.storage.StoreUser(newUser)
		if err != nil {
			returnAjaxError(writer, err.Error(), 500)
			return
		}
		newUser.SetId(userId)
		sessExpireTime := time.Now().Add(time.Hour * 24)
		newSess := dto.NewSession(
			uuid.New().String(),
			userId,
			sessExpireTime.Unix(),
		)
		err = h.storage.StoreSession(newSess)
		if err == nil {
			http.SetCookie(writer, &http.Cookie{
				Name:     "SESSION",
				Value:    newSess.GetToken(),
				Path:     "/",
				Domain:   request.Host,
				Expires:  sessExpireTime,
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteDefaultMode,
			})
		}
		writer.Header().Set("Location", "/")
		writer.WriteHeader(302)
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

	err = ts.Execute(writer, data)
	if err != nil {
		http.Error(writer, "Internal Server Error", 500)
	}
}

func returnAjaxError(writer http.ResponseWriter, msg string, code int) {
	writer.Header().Set("Content-Type", "text/html")
	http.Error(writer, msg, code)
}
