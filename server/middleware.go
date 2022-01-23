package server

import "net/http"

func checkAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if false {
			http.Error(w, "Bad auth!", 403)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
