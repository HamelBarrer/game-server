package model

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type Claim struct {
	User string `json:"user"`
	jwt.StandardClaims
}
