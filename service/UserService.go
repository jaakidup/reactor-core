package service

import (
	"fmt"

	"github.com/jaakidup/reactor-core/model"
	"github.com/jaakidup/reactor-core/store"
)

func init() {
	fmt.Println("UserService init")
}

// UserService ...
type UserService struct {
	Store *store.Store
}

// CreateUser ...
// @user model.User
// @return error
func (us UserService) CreateUser(user model.User) (string, error) {

	id, err := us.Store.User.Save(user)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetUser ....
func (us UserService) GetUser(id string) (model.User, error) {
	storedUser, err := us.Store.User.Get(id)
	if err != nil {
		return storedUser, err
	}
	return storedUser, nil
}

// ListUsers ...
func (us UserService) ListUsers() ([]model.User, error) {
	users, err := us.Store.User.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
