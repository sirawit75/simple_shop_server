package model

import (
	"time"
)

type Logger struct {
	Username       string    `bson:"username"`
	LoginTimestamp time.Time `bson:"login_timestamp"`
}
