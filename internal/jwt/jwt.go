package jwt

import (
	"crypto/rsa"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"gitlab.com/HamelBarrer/game-server/internal/model"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := os.ReadFile("internal/security/private.rsa")
	if err != nil {
		log.Fatal("not readfile private")
		return
	}
	publicBytes, err := os.ReadFile("internal/security/public.rsa.pub")
	if err != nil {
		log.Fatal("not readfile public")
		return
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("not parser private key")
		return
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("not parser public key")
		return
	}
}

func CreateToken(u model.User) (string, error) {
	claims := model.Claim{
		User: u.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(3600),
			Issuer:    "access app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ValidationToken(token string) error {
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return errors.New("format incorrect")
	}

	tok, err := jwt.ParseWithClaims(token, model.Claim{}, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return err
	}

	if tok.Valid {
		return nil
	} else {
		return err
	}
}
