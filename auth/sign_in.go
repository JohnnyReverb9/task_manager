package auth

import (
	"errors"
	"task_maker/model"
)

func SignIn(login, password string) (*model.User, error) {

	u, flag := model.ValidateUser(login, password)

	if !flag {
		return nil, errors.New("invalid user")
	}

	return u, nil
}
