package middleware

import "net/http"

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Token Required.", http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}
