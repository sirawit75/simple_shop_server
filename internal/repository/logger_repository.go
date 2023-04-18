package repository

import (
	"context"
	"sirawit/shop/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoggerQuery interface {
	InsertLoginTimestamp(input model.Logger) error
}

type loggerQuery struct {
	db *mongo.Client
}

func NewLoggerQuery(db *mongo.Client) *loggerQuery {
	return &loggerQuery{db: db}
}

func ConnectToLoggerDB(dsn string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	return client, err
}
