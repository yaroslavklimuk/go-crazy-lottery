package http_handler

import (
	"github.com/yaroslavklimuk/crazy-lottery/dto"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	rawReward struct {
		Type string                 `json:"type"`
		X    map[string]interface{} `json:"-"`
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
		h.createSessionAndRedirect(writer, user.GetId(), sessExpireTime)
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
		form := request.PostForm
		user, err := h.storage.GetUserByName(form.Get("name"))
		if err != nil {
			renderError(writer, err.Error(), 500)
			return
		}
		sessExpireTime := time.Now().Add(time.Hour * 24)
		h.createSessionAndRedirect(writer, user.GetId(), sessExpireTime)
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
		cookie, _ := request.Cookie("SESSION")
		if cookie != nil {
			session, err := h.storage.GetSession(cookie.Value)
			if err != nil {
				http.Error(writer, err.Error(), 500)
				return
			}
			if session.GetExpiredAt() > time.Now().Unix() {
				user, err := h.storage.GetUserById(session.GetUserId())
				if err != nil {
					http.Error(writer, err.Error(), 500)
					return
				}

				rand.Seed(time.Now().Unix())
				rewardType := rand.Intn(3)
				if rewardType == 0 {
					var maxMoney int64
					money, err := h.storage.GetUserMoneyRewards(session.GetUserId())
					if err != nil {
						renderError(writer, err.Error(), 500)
						return
					}
					envMax, ok := os.LookupEnv("MAX_MONEY")
					if !ok {
						maxMoney = dto.MaxMoney
					} else {
						maxMoney, err = strconv.ParseInt(envMax, 10, 64)
					}
					moneyLimit := maxMoney - money
					if moneyLimit > 0 {
						reward := dto.NewMoneyReward(user.GetId(), dto.GenerateMoneyAmount(moneyLimit), false, 0)
						writer.Header().Set("Content-Type", "application/json")
						writer.WriteHeader(200)
						_, err = writer.Write([]byte(reward.Serialize()))
						if err != nil {
							http.Error(writer, err.Error(), 500)
						}
						return
					}
				}
				if rewardType == 1 {
					var maxItems int64
					items, err := h.storage.GetUserItemRewards(session.GetUserId())
					if err != nil {
						renderError(writer, err.Error(), 500)
						return
					}
					envMax, ok := os.LookupEnv("MAX_ITEMS")
					if !ok {
						maxItems = dto.MaxItems
					} else {
						maxItems, err = strconv.ParseInt(envMax, 10, 64)
					}
					itemsLimit := maxItems - items
					if itemsLimit > 0 {
						reward := dto.NewItemReward(user.GetId(), dto.GenerateItemType(), false, 0)
						writer.Header().Set("Content-Type", "application/json")
						writer.WriteHeader(200)
						_, err = writer.Write([]byte(reward.Serialize()))
						if err != nil {
							http.Error(writer, err.Error(), 500)
						}
						return
					}
				}

				reward := dto.NewBonusReward(user.GetId(), dto.GenerateBonusAmount(500))
				writer.Header().Set("Content-Type", "application/json")
				_, err = writer.Write([]byte(reward.Serialize()))
				if err != nil {
					http.Error(writer, err.Error(), 500)
				}

				return
			}
		}
		http.Error(writer, "Unauthorized access", 403)
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
		cookie, _ := request.Cookie("SESSION")
		if cookie != nil {
			session, err := h.storage.GetSession(cookie.Value)
			if err != nil {
				http.Error(writer, err.Error(), 500)
				return
			}
			if session.GetExpiredAt() > time.Now().Unix() {
				rewardType := request.FormValue("type")
				rawRewardAmount := request.FormValue("amount")

				writer.WriteHeader(200)
				_, err = writer.Write([]byte(rewardType + " " + rawRewardAmount))
				if err != nil {
					http.Error(writer, err.Error(), 500)
				}

				if rewardType == dto.MoneyRewardType {
					rewardAmount, err := strconv.ParseInt(rawRewardAmount, 10, 16)
					if err != nil {
						http.Error(writer, err.Error(), 400)
						return
					}
					reward := dto.NewMoneyReward(session.GetUserId(), rewardAmount, false, 0)
					_, err = h.storage.StoreUserMoneyReward(reward)
					if err != nil {
						http.Error(writer, err.Error(), 500)
						return
					}
				} else if rewardType == dto.PhysItemRewardType {
					rewardKind := dto.ItemRewardType(rawRewardAmount)
					if rewardKind != dto.BicycleItem && rewardKind != dto.CrutchItem {
						http.Error(writer, "invalid type of item reward", 400)
						return
					}
					reward := dto.NewItemReward(session.GetUserId(), rewardKind, false, 0)
					_, err = h.storage.StoreUserItemReward(reward)
					if err != nil {
						http.Error(writer, err.Error(), 500)
						return
					}
				} else if rewardType == dto.BonusRewardType {
					rewardAmount, err := strconv.Atoi(rawRewardAmount)
					if err != nil {
						http.Error(writer, err.Error(), 400)
						return
					}
					err = h.storage.UpdateBalance(session.GetUserId(), rewardAmount)
					if err != nil {
						http.Error(writer, err.Error(), 400)
						return
					}
				}
				return
			}
		}
		http.Error(writer, "Unauthorized access", 403)
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

func (h *httpHandler) createSessionAndRedirect(writer http.ResponseWriter, userId int64, sessExpireTime time.Time) {
	newSess := dto.NewSession(
		randStringBytes(20, sessExpireTime.Unix()),
		userId,
		sessExpireTime.Unix(),
	)
	err := h.storage.StoreSession(newSess)
	if err != nil {
		renderError(writer, err.Error()+" ::: "+newSess.GetToken(), 500)
		return
	}
	http.SetCookie(writer, &http.Cookie{
		Name:  "SESSION",
		Value: newSess.GetToken(),
		Path:  "/",
		//Domain:   request.Host,
		Expires:  sessExpireTime,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	})
	writer.Header().Set("Location", "/")
	writer.WriteHeader(302)
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int, t int64) string {
	b := make([]byte, n)
	rand.Seed(t)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
