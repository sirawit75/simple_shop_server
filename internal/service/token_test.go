package service

import (
	"sirawit/shop/pkg/random"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	sign := random.RandomString(32)
	username := random.RandomUsername()
	duration := time.Minute
	tokenManager := NewTokenManager(sign)
	token, err := tokenManager.CreateToken(username, duration)
	assert.NotNil(t, token)
	assert.NoError(t, err)

	payload, err := tokenManager.VerifyToken(token)
	assert.NotNil(t, payload)
	assert.NoError(t, err)
	assert.Equal(t, payload, username)

}

func TestExpiredToken(t *testing.T) {
	sign := random.RandomString(32)
	username := random.RandomUsername()
	duration := -time.Minute
	tokenManager := NewTokenManager(sign)
	token, err := tokenManager.CreateToken(username, duration)
	assert.NotNil(t, token)
	assert.NoError(t, err)

	payload, err := tokenManager.VerifyToken(token)
	assert.Error(t, err)
	assert.Equal(t, payload, "")

}

func TestInvalidToken(t *testing.T) {
	sign := random.RandomString(32)
	token := random.RandomString(32)
	tokenManager := NewTokenManager(sign)
	assert.NotNil(t, token)

	payload, err := tokenManager.VerifyToken(token)
	assert.Error(t, err)
	assert.Equal(t, payload, "")

}
