package user

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/HamelBarrer/game-server/internal/function"
	"gitlab.com/HamelBarrer/game-server/internal/jwt"
	"gitlab.com/HamelBarrer/game-server/internal/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var u model.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
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

	sendToken := model.ResponseToken{Token: token}

	json.NewEncoder(w).Encode(sendToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(3600 * time.Second),
	})
}
