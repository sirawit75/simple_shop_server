package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendWelcomeEmail = "task:send_welcome_email"

type SendWelcomeEmailPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (t *redisTaskDistributor) DistributeTaskSendWelcomeEmail(ctx context.Context, payload SendWelcomeEmailPayload, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("faild to mashal task payload: %v", err)
	}
	task := asynq.NewTask(TaskSendWelcomeEmail, jsonPayload, opts...)
	info, err := t.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %v", err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueue task")
	return nil
}

func (p *redisTaskProcessor) ProcessTaskSendWelcomeEmail(ctx context.Context, task *asynq.Task) error {
	var payload SendWelcomeEmailPayload
	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}
	subject := "Welcome to Sirawit's shop!"
	content := fmt.Sprintf(`
		<h1>Hello %v</h1>
		<p>Thank you for registering</p>
	`, payload.Username)
	to := []string{payload.Email}
	err = p.mailer.SendEmail(subject, content, to)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	// client := p.db.Database(DB).Collection(LoginTimestamp)
	// if _, err := insertLoginTimestamp(client, payload.Username); err != nil {
	// 	return err
	// }
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", payload.Email).Msg("processed all tasks")
	return nil
}
