package session

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ppzxc/chattools/domain"
)

type Adapter interface {
	Login(sessionId string, userId int64, deviceId string) error
	Logout(sessionId string) error
	GetSession(sessionId string) (domain.SessionAdapter, bool)
	GetSessionByUserId(userId int64) (map[string]domain.SessionAdapter, error)
	Register(session domain.SessionAdapter) error
	Unregister(sessionId string)

	Subscribe(ctx context.Context, key string) (*redis.PubSub, error)
	Publish(ctx context.Context, key string, message interface{}) error
}
