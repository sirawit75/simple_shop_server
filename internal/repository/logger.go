package repository

import (
	"context"
	"sirawit/shop/internal/model"
	"time"
)

const (
	DB             = "logs"
	LoginTimestamp = "login_timestamp"
	// EmailStatus    = "email_status"
)

func (l *loggerQuery) InsertLoginTimestamp(input model.Logger) error {
	client := l.db.Database(DB).Collection(LoginTimestamp)
	_, err := client.InsertOne(context.Background(), model.Logger{
		Username:       input.Username,
		LoginTimestamp: time.Now(),
	})
	return err

}
