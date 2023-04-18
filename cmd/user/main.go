package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sirawit/shop/internal/app"
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/repository"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"
	"sirawit/shop/task"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	// livenessProbe

	_, err := os.Create("/tmp/user")
	if err != nil {
		log.Fatal().Err(err)
	}
	defer os.Remove("/tmp/user")

	// load config

	config, err := config.LoadUserConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	// intdb

	db, err := repository.ConnectToUserDB(config.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	log.Info().Msg("connect to user db!")

	//grpc client

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, config.GrpcLoggerServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Err(err).Msg("cannot connet to worker service")
	}
	defer conn.Close()
	log.Info().Msgf("start grpc client(connect to logger service) at %v", config.GrpcLoggerServerAddress)

	// setup server&service
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddr,
	}

	taskDistributor := task.NewRedisTaskDistributor(redisOpt)

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo, config)
	userServer := app.NewUserServer(userSvc, conn, taskDistributor)

	// server options

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// setup server
	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userServer)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot reigster server")
	}
	mux := http.NewServeMux()
	withCors := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)
	mux.Handle("/", grpcMux)
	mux.HandleFunc("/user/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := app.HttpLogger(withCors)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	srv := &http.Server{
		Addr:    config.HttpServerAddress,
		Handler: handler,
	}
	//start server
	go func() {
		log.Info().Msgf("start http server at %v", config.HttpServerAddress)
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("cannot start  server")
		}
	}()
	<-ctx.Done()
	stop()
	log.Info().Msg("shutting down gracefully, press Ctrl+c again to force")
	timeOutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(timeOutCtx); err != nil {
		log.Err(err)
	}
}
