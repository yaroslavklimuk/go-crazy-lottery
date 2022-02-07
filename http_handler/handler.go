package http_handler

import (
	"github.com/google/uuid"
	"github.com/yaroslavklimuk/crazy-lottery/dto"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"html/template"
	"net/http"
	"time"
)

type (
	httpHandler struct {
		storage storage.Storage
	}
	mainInfo struct {
		Balance     int64
		ItemsCount  int64
		MoneyReward int64
	}
)

func (h *httpHandler) Index(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		cookie, _ := request.Cookie("SESSION")
		if cookie != nil {
			session, err := h.storage.GetSession(cookie.Value)
			if err != nil {
				renderError(writer, err.Error(), 500)
				return
			}
			if session.GetExpiredAt() > time.Now().Unix() {
				user, err := h.storage.GetUserById(session.GetUserId())
				if err != nil {
					renderError(writer, err.Error(), 500)
					return
				}

				money, err := h.storage.GetUserMoneyRewards(session.GetUserId())
				if err != nil {
					renderError(writer, err.Error(), 500)
					return
				}

				items, err := h.storage.GetUserItemRewards(session.GetUserId())
				if err != nil {
					renderError(writer, err.Error(), 500)
					return
				}

				templates := []string{
					"./templates/index.html",
					"./templates/base.html",
				}
				renderTemplate(templates, writer, mainInfo{
					Balance:     user.GetBalance(),
					ItemsCount:  items,
					MoneyReward: money,
				})
				return
			}
		}
		writer.Header().Set("Location", "/login")
		writer.WriteHeader(302)
	default:
		http.Error(writer, "Sorry, only GET requests are supported.", 405)
		return
	}
}

func (h *httpHandler) Register(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/register" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "GET":
		cookie, _ := request.Cookie("SESSION")
		if cookie != nil {
			ok, err := h.checkSession(cookie.Value, time.Now().Unix())
			if err != nil {
				renderError(writer, err.Error(), 500)
				return
			}
			if ok == true {
				writer.Header().Set("Location", "/")
				writer.WriteHeader(302)
				return
			}
		}

		templates := []string{
			"./templates/register.html",
			"./templates/base.html",
		}
		renderTemplate(templates, writer, nil)
	case "POST":
		form := request.PostForm
		user, err := h.storage.GetUserByName(form.Get("name"))
		if err != nil {
			renderError(writer, err.Error(), 500)
			return
		}
		if user != nil {
			renderError(writer, "User already exists", 400)
			return
		}
		newUser := dto.NewUser(
			0,
			form.Get("name"),
			form.Get("banc_acc"),
			form.Get("address"),
			0,
			form.Get("passwd"),
		)
		userId, err := h.storage.StoreUser(newUser)
		if err != nil {
			renderError(writer, err.Error(), 500)
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

func (h *httpHandler) Login(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/login" {
		http.Error(writer, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "GET":
		cookie, _ := request.Cookie("SESSION")
		if cookie != nil {
			ok, err := h.checkSession(cookie.Value, time.Now().Unix())
			if err != nil {
				renderError(writer, err.Error(), 500)
				return
			}
			if ok == true {
				writer.Header().Set("Location", "/")
				writer.WriteHeader(302)
				return
			}
		}

		templates := []string{
			"./templates/login.html",
			"./templates/base.html",
		}
		renderTemplate(templates, writer, nil)
	case "POST":
		writer.Header().Set("Content-Type", "application/json")
	default:
		http.Error(writer, "Sorry, only GET or POST requests are supported.", 405)
		return
	}
}

func (h *httpHandler) GetReward(writer http.ResponseWriter, request *http.Request) {
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

func (h *httpHandler) SubmitReward(writer http.ResponseWriter, request *http.Request) {
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

func (h *httpHandler) checkSession(token string, timestamp int64) (bool, error) {
	session, err := h.storage.GetSession(token)
	if err != nil {
		return false, err
	}
	ok := session.GetExpiredAt() > timestamp
	return ok, nil
}

func MakeHttpHandler(st storage.Storage) *httpHandler {
	return &httpHandler{storage: st}
}

func renderTemplate(templates []string, writer http.ResponseWriter, data interface{}) {
	ts, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	err = ts.Execute(writer, data)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}
}

func renderError(writer http.ResponseWriter, msg string, code int) {
	writer.Header().Set("Content-Type", "text/html")
	http.Error(writer, msg, code)
}
