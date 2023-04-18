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
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	// livenessProbe

	_, err := os.Create("/tmp/product")
	if err != nil {
		log.Fatal().Err(err)
	}
	defer os.Remove("/tmp/product")

	// load config

	config, err := config.LoadProductConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	// intdb

	db, err := repository.ConnectToProductDB(config.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	log.Info().Msg("connect to product db!")

	//setup service && server

	productQuery := repository.NewProductQuery(db)
	productService := service.NewProductService(productQuery, config)
	server := app.NewProductServer(productService, config)

	//json options

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterProductServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register server")
	}
	mux := http.NewServeMux()
	withCors := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)
	mux.Handle("/", grpcMux)
	mux.HandleFunc("/product/healthz", func(w http.ResponseWriter, req *http.Request) {
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
	// go func(){
	// 	if err := s.ListenAndServe(); err != nil && err!=http.ErrServerClosed{
	// 		log.Fatalf("listen: %s\n", err);
	// 	}
	// }();
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
