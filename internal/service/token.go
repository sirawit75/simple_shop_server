package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (any, error)
}

type JWTManager struct {
	sign string
}

func NewTokenManager(sign string) TokenManager {
	return &JWTManager{sign}
}

func (m *JWTManager) CreateToken(username string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(duration).Unix(),
		Audience:  username,
	})
	ss, err := token.SignedString([]byte(m.sign))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (m *JWTManager) VerifyToken(token string) (any, error) {
	decodedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signin method")
		}
		return []byte(m.sign), nil
	})
	if err != nil {
		return "", err
	}
	if cliam, ok := decodedToken.Claims.(jwt.MapClaims); ok {
		aud := cliam["aud"]
		return aud, nil
	}
	return "", err

}
