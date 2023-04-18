package service

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *userService) Login(username, password string) (*UserRes, error) {
	err := ValidateUsername(username)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = ValidatePassword(password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, err := u.db.FindUserByUsername(username)
	if err != nil {
		if strings.Contains(err.Error(), NotFoundStatus) {
			return nil, status.Error(codes.NotFound, UserNotFound)
		}
		return nil, status.Errorf(codes.Internal, "failed to find user %v", err)

	}
	if ok := CheckPasswordHash(password, user.Password); !ok {
		return nil, status.Error(codes.NotFound, "incorrect password")
	}
	token, err := u.tokenManager.CreateToken(user.Username, u.config.TokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create token %v", err)
	}
	return &UserRes{
		User:  *user,
		Token: token,
	}, nil
}
