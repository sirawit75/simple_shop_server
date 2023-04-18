package service

import (
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/repository"
)

type LoggerService interface {
	InsertLoginTimestamp(input model.Logger) (*model.Logger, error)
}

type loggerService struct {
	db repository.LoggerQuery
}

func NewLoggerService(db repository.LoggerQuery) LoggerService {
	return &loggerService{db}
}
