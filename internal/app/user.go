package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"
	"sirawit/shop/task"
	"time"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserToUserRes(result *service.UserRes) *pb.RegisterRes {
	return &pb.RegisterRes{
		User: &pb.User{
			ID:        result.User.ID,
			Username:  result.User.Username,
			Email:     result.User.Email,
			CreatedAt: timestamppb.New(result.User.CreatedAt),
		},
		Token: result.Token,
	}
}

func (u *userServer) TestApi(context.Context, *emptypb.Empty) (*pb.TestApiRes, error) {
	return &pb.TestApiRes{
		Message: "d krub oat",
	}, nil
}

func (u *userServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	result, err := u.userService.Register(model.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	opts := []asynq.Option{
		asynq.Timeout(10 * time.Second),
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
	}
	err = u.taskDistributor.DistributeTaskSendWelcomeEmail(
		ctx, task.SendWelcomeEmailPayload{
			Username: req.GetUsername(),
			Email:    req.GetEmail(),
		}, opts...,
	)
	if err != nil {
		log.Err(err).Msg("Send to worker service failed")
	} else {
		log.Info().Msg("Send to worker service success")
	}
	_, err = u.loggerClient.SendLoginTimestampToLogger(ctx, &pb.LoginTimestamp{
		Username: result.User.Username,
	})
	if err != nil {
		log.Err(err).Msg("Send to logger service failed")
	} else {
		log.Info().Msg("Send to logger service success")
	}
	return convertUserToUserRes(result), nil
}


func (u *userServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	result, err := u.userService.Login(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return (*pb.LoginRes)(convertUserToUserRes(result)), nil
}
