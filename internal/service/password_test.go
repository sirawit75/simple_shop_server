package service

import (
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := random.RandomString(6)
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotNil(t, hashedPassword)
	assert.NotEqual(t, hashedPassword, password)
}

func TestCheckPassword(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		password := random.RandomString(6)
		hashedPassword, err := HashPassword(password)
		assert.NoError(t, err)
		assert.NotNil(t, hashedPassword)
		assert.NotEqual(t, hashedPassword, password)

		result := CheckPasswordHash(password, hashedPassword)
		assert.True(t, result)
	})

	t.Run("failed", func(t *testing.T) {
		password := random.RandomString(6)
		hashedPassword, err := HashPassword(password)
		assert.NoError(t, err)
		assert.NotNil(t, hashedPassword)
		assert.NotEqual(t, hashedPassword, password)

		result := CheckPasswordHash(random.RandomString(6), hashedPassword)
		assert.False(t, result)
	})
}
