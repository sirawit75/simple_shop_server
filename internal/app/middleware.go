package app

import (
	"context"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(res, req)
		logger := log.Info()
		logger.Str("protocol", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Msg("received a HTTP request")
	})
}

func (m *productServer) getUsernameFromToken(ctx context.Context) (string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Debug().Msgf("md = %v", md)
	token := md.Get("Authorization")
	log.Debug().Msgf("token = %v", token)
	if token == nil {
		return "", status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	accessToken := strings.TrimPrefix(token[0], "Bearer")
	accessToken = strings.ReplaceAll(accessToken, " ", "")
	log.Debug().Msgf("acc = %v", accessToken)
	username, err := m.tokenManager.VerifyToken(accessToken)
	if err != nil {
		return "", err
	}
	usr, ok := username.(string)
	log.Debug().Msgf("usr = %v", usr)
	if !ok {
		return "", status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	return usr, nil

}
