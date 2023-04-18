package service

import (
	"sirawit/shop/internal/model"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRes struct {
	User  model.User
	Token string
}

const (
	UsernameAlreadyExists = "username already exists"
	SQLSTATE23505         = "SQLSTATE 23505"
	EmailAlreadyExists    = "email already exists"
	NotFoundStatus        = "record not found"
	UserNotFound          = "user not found"
	ProductNotFound       = "product not found"
)

func (u *userService) Register(input model.User) (*UserRes, error) {
	err := ValidateUsername(input.Username)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = ValidatePassword(input.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = ValidateEmail(input.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	input.Password, err = HashPassword(input.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password %v", err)
	}
	user, err := u.db.Register(input)
	if err != nil {
		if strings.Contains(err.Error(), SQLSTATE23505) {
			if strings.Contains(err.Error(), "username") {
				return nil, status.Error(codes.AlreadyExists, UsernameAlreadyExists)
			}
			return nil, status.Error(codes.AlreadyExists, EmailAlreadyExists)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %v", err)

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
