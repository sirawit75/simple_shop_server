package repository

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testDB *gorm.DB
var testUserQuery UserQuery

func TestConnectToUserDB(t *testing.T) {
	var err error
	config, err := config.LoadUserConfig("../../cmd/user")
	assert.NoError(t, err)
	testDB, err = ConnectToUserDB(config.DSN)
	assert.NoError(t, err)
	testUserQuery = NewUserRepository(testDB)

}
func TestRegister(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		input := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomUsername(),
			Email:    random.RandomEmail(),
		}
		user, err := testUserQuery.Register(input)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user.Username, input.Username)
		assert.Equal(t, user.Email, input.Email)
		assert.NotNil(t, user.CreatedAt)
		assert.Greater(t, user.ID, uint64(0))
	})
	t.Run("failed (user alreaydy exists)", func(t *testing.T) {
		input := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomUsername(),
			Email:    random.RandomEmail(),
		}
		_, _ = testUserQuery.Register(input)
		user2, err := testUserQuery.Register(input)
		assert.Empty(t, user2)
		assert.Error(t, err)
	})

	t.Run("failed (email alreaydy exists)", func(t *testing.T) {
		input := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomUsername(),
			Email:    random.RandomEmail(),
		}
		_, _ = testUserQuery.Register(input)
		input2 := input
		input2.Username = random.RandomUsername()
		user2, err := testUserQuery.Register(input2)
		assert.Empty(t, user2)
		assert.Error(t, err)
	})
}

func TestFindUserByUsername(t *testing.T) {
	t.Run("failed", func(t *testing.T) {
		username := random.RandomUsername()
		result, err := testUserQuery.FindUserByUsername(username)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("pass", func(t *testing.T) {
		input := model.User{
			Username: random.RandomUsername(),
			Password: random.RandomUsername(),
			Email:    random.RandomEmail(),
		}
		user, err := testUserQuery.Register(input)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		result, err := testUserQuery.FindUserByUsername(user.Username)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, user.Email, result.Email)
		assert.Equal(t, user.Username, result.Username)
	})
}
