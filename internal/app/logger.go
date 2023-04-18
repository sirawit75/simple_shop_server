package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/pb"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *loggerServer) CheckRediness(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	// perform any necessary checks to determine whether the service is ready
	// return an error if the service is not ready
	if !serviceIsReady() {
		return nil, status.Errorf(codes.Unavailable, "service is not ready")
	}

	return &emptypb.Empty{}, nil
}

func (l *loggerServer) SendLoginTimestampToLogger(ctx context.Context, req *pb.LoginTimestamp) (*pb.LoginTimestamp, error) {
	result, err := l.service.InsertLoginTimestamp(model.Logger{
		Username: req.GetUsername(),
	})
	if err != nil {
		log.Err(err).Msg("cannot insert to logger db")
		return nil, err
	}
	log.Info().Msg("insert into logger db success")
	return &pb.LoginTimestamp{
		Username:       result.Username,
		LoginTimestamp: timestamppb.New(result.LoginTimestamp),
	}, nil
}
