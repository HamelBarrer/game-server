package function

import (
	"context"

	"gitlab.com/HamelBarrer/game-server/internal/model"
	"gitlab.com/HamelBarrer/game-server/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidationUser(email string) (model.User, bool) {
	u := &model.User{}
	collection := storage.Connection().Database("testing").Collection("users")

	if err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(u); err != nil {
		return *u, false
	}

	return *u, true
}
