package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sirawit/shop/internal/config"
	"sirawit/shop/mail"
	"sirawit/shop/task"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func main() {

	// livenessProbe

	_, err := os.Create("/tmp/worker")
	if err != nil {
		log.Fatal().Err(err)
	}
	defer os.Remove("/tmp/worker")

	// load config

	config, err := config.LoadWorkerConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddr,
	}

	//redis task processor
	// go func() {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := task.NewRedisTaskProcessor(redisOpt, mailer)
	go func() {
		err = taskProcessor.Start(config.RedisAddr)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to start task processor")
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// create a readiness endpoint
	http.HandleFunc("/worker/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// fmt.Fprintf(w, "OK")
	})
	go func() {
		// log.Info().Msgf("start http server at %v", config.HttpServerAddress)
		if err = http.ListenAndServe(config.WorkerHealthz, nil); err != nil && err != http.ErrServerClosed {
			log.Panic().Err(err)
		}
	}()
	<-ctx.Done()
	stop()
	taskProcessor.Shutdown()
	// timeOutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(timeOutCtx); err != nil {
	// 	log.Err(err)
	// }

}
