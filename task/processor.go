package task

import (
	"context"
	"sirawit/shop/mail"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type TaskProcessor interface {
	Start(addr string) error
	// ProcessTaskInsetLoginTimestamp(ctx context.Context, task *asynq.Task) error
	ProcessTaskSendWelcomeEmail(ctx context.Context, task *asynq.Task) error
	Shutdown()
}

type redisTaskProcessor struct {
	server *asynq.Server
	// db     *mongo.Client
	mailer mail.EmailSender
}

func NewRedisTaskProcessor(
	redisOpt asynq.RedisClientOpt,
	// db *mongo.Client,
	mailer mail.EmailSender,
) TaskProcessor {
	server := asynq.NewServer(redisOpt, asynq.Config{
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Err(err).Str("type", task.Type()).
				Bytes("payload", task.Payload()).Msg("process task failed")
		}),
	})
	return &redisTaskProcessor{
		server: server,
		// db:     db,
		mailer: mailer,
	}
}

func (p *redisTaskProcessor) Start(addr string) error {
	mux := asynq.NewServeMux()
	// mux.HandleFunc(TaskInsertLoginTimestamp, p.ProcessTaskInsetLoginTimestamp)
	mux.HandleFunc(TaskSendWelcomeEmail, p.ProcessTaskSendWelcomeEmail)
	log.Info().Msgf("start redis server at %v", addr)
	return p.server.Run(mux)
}

func (p *redisTaskProcessor) Shutdown() {
	log.Info().Msg("shutting down gracefully")
	p.server.Shutdown()
}
