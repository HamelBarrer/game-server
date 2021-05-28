package function

import (
	"errors"

	"gitlab.com/HamelBarrer/game-server/internal/model"
	"gitlab.com/HamelBarrer/game-server/internal/security"
)

func VerificationLogin(email, password string) (model.User, error) {
	user, exist := ValidationUser(email)
	if !exist {
		return model.User{}, errors.New("user not found")
	}

	equals, err := security.CompareHash(password, user.Password)
	if err != nil {
		return model.User{}, err
	}

	if !equals {
		return model.User{}, errors.New("user o password incorrect")
	}

	return user, nil
}
