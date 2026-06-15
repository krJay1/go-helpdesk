package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request Received from:", r.URL.Path)
		log.Println("Request Received from:", r)
		next.ServeHTTP(w, r)
	})

}
