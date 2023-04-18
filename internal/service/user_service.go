package service

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/repository"
)

type UserService interface {
	Register(input model.User) (*UserRes, error)
	Login(username, password string) (*UserRes, error)
}

type userService struct {
	db           repository.UserQuery
	tokenManager TokenManager
	config       config.UserConfig
}

func NewUserService(db repository.UserQuery, config config.UserConfig) UserService {
	tokenManager := NewTokenManager(config.Sign)
	return &userService{db: db, config: config, tokenManager: tokenManager}
}
