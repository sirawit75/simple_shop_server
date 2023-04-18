package service

import (
	"sirawit/shop/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *loggerService) InsertLoginTimestamp(input model.Logger) (*model.Logger, error) {
	err := l.db.InsertLoginTimestamp(input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert into login timestamp %v", err)
	}
	return &input, nil
}
