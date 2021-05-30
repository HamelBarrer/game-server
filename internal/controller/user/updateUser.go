package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"gitlab.com/HamelBarrer/game-server/internal/model"
	"gitlab.com/HamelBarrer/game-server/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := &model.User{}

	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := storage.Connection().Database("testing").Collection("users")
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{"$set", bson.D{
			{"firstName", u.FirstName},
			{"lastName", u.LastName},
			{"username", u.Username},
			{"date_birth", u.DateBirth},
			{"updated_at", time.Now()},
		}},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
