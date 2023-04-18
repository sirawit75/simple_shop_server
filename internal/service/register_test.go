package service

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/mock"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		user := model.User{
			Username: "pass123",
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}

		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Register(user)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.User.ID)
		assert.NotNil(t, result.Token)
		assert.Equal(t, result.User.Username, user.Username)
		assert.Equal(t, result.User.Email, user.Email)
		assert.NotEqual(t, result.User.Password, user.Password)
	})

	t.Run("failed (server error)", func(t *testing.T) {
		user := model.User{
			Username: bobby123,
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}

		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Internal")
	})

	t.Run("failed (username already exist)", func(t *testing.T) {
		user := model.User{
			Username: "username",
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}

		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), UsernameAlreadyExists)
	})

	t.Run("failed (email already exist)", func(t *testing.T) {
		user := model.User{
			Username: "email123",
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}

		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Register(user)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), EmailAlreadyExists)
	})

	t.Run("invalid username", func(t *testing.T) {
		user := model.User{
			Username: "Afd23_",
			Password: random.RandomString(6),
			Email:    random.RandomEmail(),
		}
		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Register(user)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "username must contain only lowercase letters, digits, or underscore")
	})

	t.Run("invalid password", func(t *testing.T) {
		user := model.User{
			Username: bobby123,
			Password: random.RandomString(3),
			Email:    random.RandomEmail(),
		}
		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Register(user)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password must contain from 6-100 characters")
	})
	t.Run("email", func(t *testing.T) {
		user := model.User{
			Username: bobby123,
			Password: random.RandomString(6),
			Email:    random.RandomUsername(),
		}
		testDB := mock.NewUserRepositoryMock()
		testUserService := NewUserService(testDB, config.UserConfig{})
		result, err := testUserService.Register(user)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "is not a valid email")
	})
}
