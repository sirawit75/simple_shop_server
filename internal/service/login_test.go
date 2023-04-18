package service

import (
	"errors"
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/mock"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

const bobby123 = "bobby123"

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := model.User{
			Username: bobby123,
			Password: random.RandomString(6),
		}
		hashedPassword, err := HashPassword(user.Password)
		assert.NoError(t, err)
		testDB := mock.NewUserRepositoryMock()
		testDB.On("FindUserByUsername", user.Username).Return(&model.User{Username: user.Username, Password: hashedPassword}, nil)
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Login(user.Username, user.Password)
		assert.NotNil(t, result)
		assert.NoError(t, err)
		assert.NotEqual(t, user.Username, result.User.Password)
	})
	t.Run("password not match", func(t *testing.T) {
		username := bobby123
		password := random.RandomString(6)
		testDB := mock.NewUserRepositoryMock()
		testDB.On("FindUserByUsername", username).Return(&model.User{}, nil)
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Login(username, password)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "incorrect password")
	})

	t.Run("user not found", func(t *testing.T) {
		username := bobby123
		password := random.RandomString(6)
		testDB := mock.NewUserRepositoryMock()
		testDB.On("FindUserByUsername", username).Return(&model.User{}, errors.New(NotFoundStatus))
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Login(username, password)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), UserNotFound)
	})

	t.Run("invalid password", func(t *testing.T) {

		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		username := bobby123
		password := random.RandomString(3)
		result, err := testUserService.Login(username, password)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password must contain from 6-100 characters")
	})

	t.Run("invalid username", func(t *testing.T) {

		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		username := "ASDsd2"
		password := random.RandomString(6)
		result, err := testUserService.Login(username, password)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "username must contain only lowercase letters, digits, or underscore")
	})
}
