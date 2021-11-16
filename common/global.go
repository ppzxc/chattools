package common

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var Validate *validator.Validate

var (
	WebSocketWriteTimeOut = 5 * time.Second
	QueryTimeOut          = 5 * time.Second
	RedisCmdTimeOut       = 5 * time.Second
	GrpcRequestTimeOut    = 5 * time.Second

	FindCount int64 = 100
)
