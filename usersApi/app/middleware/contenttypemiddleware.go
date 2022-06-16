package middleware

import (
	"net/http"
	"usersApi/mediatypenames"
)

func SetJsonContentMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mediatypenames.ApplicationJson)
		h.ServeHTTP(w, r)
	})

}
