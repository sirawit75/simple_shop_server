package app

import (
	"context"
	"sirawit/shop/pkg/pb"
	"sirawit/shop/task"
	"time"

	"github.com/hibiken/asynq"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceIsReady() bool {
	// perform any necessary checks to determine whether the service is ready
	// return true if the service is ready, false otherwise
	return true
}

// func (w *workerServer) CheckRediness(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
// 	// perform any necessary checks to determine whether the service is ready
// 	// return an error if the service is not ready
// 	if !serviceIsReady() {
// 		return nil, status.Errorf(codes.Unavailable, "service is not ready")
// 	}

// 	return &emptypb.Empty{}, nil
// }

// func (w *workerServer) InsertLoginTimestamp(ctx context.Context, req *pb.LoginTimestampReq) (*emptypb.Empty, error) {
// 	opts := []asynq.Option{
// 		asynq.Timeout(10 * time.Second),
// 		asynq.MaxRetry(10),
// 		asynq.ProcessIn(10 * time.Second),
// 	}
// 	return &emptypb.Empty{}, w.taskDistributor.DistributeTaskInsertLoginTimestamp(ctx, task.InsertLoginTimeStampPayload{
// 		Username: req.GetUsername(),
// 	}, opts...)
// }

func (w *workerServer) SendWelcomeEmail(ctx context.Context, req *pb.SendWelcomeEmailReq) (*emptypb.Empty, error) {
	opts := []asynq.Option{
		asynq.Timeout(10 * time.Second),
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
	}
	return &emptypb.Empty{}, w.taskDistributor.DistributeTaskSendWelcomeEmail(
		ctx, task.SendWelcomeEmailPayload{
			Username: req.GetUsername(),
			Email:    req.GetEmail(),
		}, opts...,
	)
}
