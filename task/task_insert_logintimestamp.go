package task

// import (
// 	"context"
// 	"fmt"
// 	"sirawit/shop/internal/model"
// 	"time"

// 	"github.com/hibiken/asynq"
// 	"github.com/rs/zerolog/log"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type InsertLoginTimeStampPayload struct {
// 	Username string
// }

// const (
// 	DB                       = "logs"
// 	LoginTimestamp           = "login_timestamp"
// 	TaskInsertLoginTimestamp = "task:insert_login_timestamp"
// )

// func (t *redisTaskDistributor) DistributeTaskInsertLoginTimestamp(ctx context.Context, payload InsertLoginTimeStampPayload, opts ...asynq.Option) error {
// 	task := asynq.NewTask(TaskInsertLoginTimestamp, []byte(payload.Username), opts...)
// 	info, err := t.client.EnqueueContext(ctx, task)
// 	if err != nil {
// 		return fmt.Errorf("failed to enqueue task: %v", err)
// 	}
// 	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
// 		Int("max_retry", info.MaxRetry).Msg("enqueue task")
// 	return nil
// }

// func insertLoginTimestamp(client *mongo.Collection, username string) (*mongo.InsertOneResult, error) {
// 	return client.InsertOne(context.Background(), model.Logger{
// 		Username:       username,
// 		LoginTimestamp: time.Now(),
// 	})
// }

// func (p *redisTaskProcessor) ProcessTaskInsetLoginTimestamp(ctx context.Context, task *asynq.Task) error {
// 	client := p.db.Database(DB).Collection(LoginTimestamp)
// 	if _, err := insertLoginTimestamp(client, string(task.Payload())); err != nil {
// 		return err
// 	}
// 	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
// 		Str("username", string(task.Payload())).Msg("processed task")
// 	return nil
// }
