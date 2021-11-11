package session

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ppzxc/chattools/domain"
)

type Adapter interface {
	Login(ctx context.Context, sessionId string, userId int64, userName string, deviceId string) error
	Logout(ctx context.Context, sessionId string) error
	GetSession(ctx context.Context, sessionId string) (domain.SessionAdapter, bool)
	GetSessionByUserId(ctx context.Context, userId int64) (map[string]domain.SessionAdapter, error)
	ExistSession(ctx context.Context, sessionId string) bool
	Register(ctx context.Context, session domain.SessionAdapter) error
	Unregister(ctx context.Context, sessionId string)

	Subscribe(ctx context.Context, key string) (*redis.PubSub, error)
	Publish(ctx context.Context, key string, message interface{}) error
}
