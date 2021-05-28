package middleware

import (
	"net/http"

	"gitlab.com/HamelBarrer/game-server/internal/jwt"
)

func VerificationToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "token not provider", http.StatusBadRequest)
			return
		}
		err := jwt.ValidationToken(token)
		if err != nil {
			http.Error(w, "token incorrect "+err.Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
