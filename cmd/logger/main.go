package main

import (
	"net"
	"os"
	"os/signal"
	"sirawit/shop/internal/app"
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/repository"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"
	"syscall"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// livenessProbe

	_, err := os.Create("/tmp/logger")
	if err != nil {
		log.Fatal().Err(err)
	}
	defer os.Remove("/tmp/logger")

	// load config

	config, err := config.LoadLoggerConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	//connect to db
	client, err := repository.ConnectToLoggerDB(config.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to logger db")
	}

	log.Info().Msg("connect to logger db!")

	//setup service & server
	loggerRepository := repository.NewLoggerQuery(client)
	loggerService := service.NewLoggerService(loggerRepository)
	server := app.NewLoggerServer(loggerService)
	if err != nil {
		log.Fatal().Msg("cannot create logger server")
	}

	// grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GrpcLoggerServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot create listener")
	}
	go func() {
		log.Info().Msgf("start gRPC server at %v", listener.Addr().String())
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal().Msg("cannot start grpc server")
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Info().Msg("Initiating graceful shutdown...")
	grpcServer.GracefulStop()
	log.Info().Msg("Server stopped.")
}
