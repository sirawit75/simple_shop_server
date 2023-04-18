package app

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/pkg/pb"
	"sirawit/shop/task"
)

type workerServer struct {
	pb.UnimplementedWorkerServiceServer
	config          config.WorkerConfig
	taskDistributor task.TaskDistributor
}

func NewWorkerServer(
	config config.WorkerConfig,
	taskDistributor task.TaskDistributor,
) *workerServer {
	return &workerServer{config: config, taskDistributor: taskDistributor}
}
