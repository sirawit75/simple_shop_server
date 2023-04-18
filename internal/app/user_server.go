package app

import (
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"
	"sirawit/shop/task"

	"google.golang.org/grpc"
)

type userServer struct {
	userService service.UserService
	pb.UnimplementedUserServiceServer
	// workerClient pb.WorkerServiceClient
	loggerClient    pb.LoggerServiceClient
	taskDistributor task.TaskDistributor
}

func NewUserServer(userService service.UserService, conn *grpc.ClientConn, taskDistributor task.TaskDistributor) *userServer {
	loggerClient := pb.NewLoggerServiceClient(conn)
	return &userServer{
		userService:     userService,
		loggerClient:    loggerClient,
		taskDistributor: taskDistributor,
	}
}
