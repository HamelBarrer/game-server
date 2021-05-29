package role

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/HamelBarrer/game-server/internal/model"
	"gitlab.com/HamelBarrer/game-server/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := storage.Connection().Database("testing").Collection("roles")
	rl := &model.Role{}

	err := json.NewDecoder(r.Body).Decode(rl)
	if err != nil {
		http.Error(w, "structure json incorrect", http.StatusBadRequest)
		return
	}

	_, err = collection.InsertOne(ctx, bson.D{
		{"name", rl.Name},
		{"state", true},
		{"created_at", time.Now()},
	})
	if err != nil {
		http.Error(w, "data insert error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
