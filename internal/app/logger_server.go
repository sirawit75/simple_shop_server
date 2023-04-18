package app

import (
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"
)

type loggerServer struct {
	pb.UnimplementedLoggerServiceServer
	service service.LoggerService
}

func NewLoggerServer(service service.LoggerService) *loggerServer {
	return &loggerServer{service: service}
}
