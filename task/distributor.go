package task

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	// DistributeTaskInsertLoginTimestamp(
	// 	ctx context.Context,
	// 	payload InsertLoginTimeStampPayload,
	// 	opts ...asynq.Option,
	// ) error
	DistributeTaskSendWelcomeEmail(ctx context.Context, payload SendWelcomeEmailPayload, opts ...asynq.Option) error
}

type redisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &redisTaskDistributor{client: client}
}
