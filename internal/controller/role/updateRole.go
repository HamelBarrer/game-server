package role

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"gitlab.com/HamelBarrer/game-server/internal/model"
	"gitlab.com/HamelBarrer/game-server/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := storage.Connection().Database("testing").Collection("roles")
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rl := &model.Role{}

	if err := json.NewDecoder(r.Body).Decode(rl); err != nil {
		http.Error(w, "structure json incorrect", http.StatusBadRequest)
		return
	}

	collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{"$set", bson.D{
			{"name", rl.Name},
			{"state", rl.State},
			{"updated_at", time.Now()},
		}},
	})
}
