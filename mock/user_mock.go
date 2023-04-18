package mock

import (
	"errors"
	"sirawit/shop/internal/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) Register(input model.User) (*model.User, error) {
	switch input.Username {
	case "pass123":
		return &input, nil
	case "username":
		return nil, errors.New("username already exists" + " " + "SQLSTATE 23505")
	case "email123":
		return nil, errors.New("email already exists" + " " + "SQLSTATE 23505")
	default:
		return nil, errors.New("")
	}

}

func (m *UserRepositoryMock) FindUserByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}
