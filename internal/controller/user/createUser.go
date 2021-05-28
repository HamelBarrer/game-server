package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/HamelBarrer/game-server/internal/function"
	"gitlab.com/HamelBarrer/game-server/internal/model"
	"gitlab.com/HamelBarrer/game-server/internal/security"
	"gitlab.com/HamelBarrer/game-server/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := storage.Connection().Database("testing").Collection("users")
	u := &model.User{}

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(w, "structure json incorrect", http.StatusBadRequest)
		return
	}

	_, exist := function.ValidationUser(u.Email)
	if exist {
		http.Error(w, "email exist", http.StatusBadRequest)
		return
	}

	u.Password, err = security.GenerateHash(u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = collection.InsertOne(ctx, bson.D{
		{"first_name", u.FirstName},
		{"last_name", u.LastName},
		{"username", u.Username},
		{"email", u.Email},
		{"password", u.Password},
		{"date_birth", u.DateBirth},
		{"created_at", time.Now()},
	})
	if err != nil {
		http.Error(w, "data insert error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
