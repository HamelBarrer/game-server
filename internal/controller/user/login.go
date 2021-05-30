package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitlab.com/HamelBarrer/game-server/internal/function"
	"gitlab.com/HamelBarrer/game-server/internal/jwt"
	"gitlab.com/HamelBarrer/game-server/internal/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u model.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(w, "format incorrect of json", http.StatusBadRequest)
		return
	}

	user, err := function.VerificationLogin(u.Email, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := jwt.CreateToken(user)
	if err != nil {
		http.Error(w, "ups "+err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(1 * time.Hour),
	})
	w.WriteHeader(http.StatusOK)
}
