package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Username  string             `bson:"username,omitempty" json:"username,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	DateBirth time.Time          `bson:"date_birth,omitempty" json:"date_birth,omitempty"`
	Avatar    string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
