package service

import (
	"github.com/jaakidup/reactor-core/store"
)

// CreateService ...
func CreateService() *Service {
	store := store.CreateStore()
	userService := UserService{Store: store}

	return &Service{
		Store: store,
		User:  userService,
	}
}

// Service ...
type Service struct {
	Store *store.Store
	User  UserService
}
