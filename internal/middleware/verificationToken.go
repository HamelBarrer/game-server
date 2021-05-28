package middleware

import (
	"net/http"

	"gitlab.com/HamelBarrer/game-server/internal/jwt"
)

func VerificationToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := jwt.ValidationToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "token incorrect", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
