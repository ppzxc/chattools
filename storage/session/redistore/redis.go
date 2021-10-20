package redistore

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ppzxc/chattools/domain"
	"github.com/ppzxc/chattools/storage/cache"
	"github.com/ppzxc/chattools/storage/session"
	"github.com/sirupsen/logrus"
)

func NewRedisSessionStore(adapter cache.Adapter) session.Adapter {
	return &redisSessionStore{
		rdb: adapter,
	}
}

type redisSessionStore struct {
	rdb cache.Adapter
}

func (r redisSessionStore) Subscribe(ctx context.Context, key string) (*redis.PubSub, error) {
	return r.rdb.Subscribe(ctx, key)
}

func (r redisSessionStore) Publish(ctx context.Context, key string, message interface{}) error {
	return r.rdb.Publish(ctx, key, message)
}

func (r redisSessionStore) Login(sessionId string, userId int64, browserId string) error {
	get, err := r.rdb.Get(sessionId)
	if err != nil || get == nil {
		return session.ErrNotRegister
	} else {
		sess := domain.NewSessionFromExternal(get)
		sess.Login(userId, browserId)
		err := r.rdb.Set(sessionId, sess.ToExternal())
		if err != nil {
			return err
		}
		return nil
	}
}

func (r redisSessionStore) Logout(sessionId string) error {
	get, err := r.rdb.Get(sessionId)
	if err != nil || get == nil {
		return session.ErrNotRegister
	} else {
		sess := domain.NewSessionFromExternal(get)
		sess.Logout()
		err := r.rdb.Set(sessionId, sess.ToExternal())
		if err != nil {
			return err
		}
		return nil
	}
}

func (r redisSessionStore) GetSession(sessionId string) (domain.SessionAdapter, bool) {
	get, err := r.rdb.Get(sessionId)
	if err != nil || get == nil {
		return nil, false
	} else {
		return domain.NewSessionFromExternal(get), true
	}
}

//TODO get user sessions in multidevce
func (r redisSessionStore) GetSessions(userId int64) (map[string]domain.SessionAdapter, bool) {
	panic("implement me")
}

func (r redisSessionStore) Register(registerSession domain.SessionAdapter) error {
	get, err := r.rdb.Get(registerSession.GetSessionId())
	if err == nil && get != nil {
		return session.ErrAlreadyRegister
	} else {
		err := r.rdb.Set(registerSession.GetSessionId(), registerSession.ToExternal())
		if err != nil {
			return err
		}
		return nil
	}
}

func (r redisSessionStore) Unregister(sessionId string) {
	err := r.rdb.Del(sessionId)
	if err != nil {
		logrus.WithError(err).Debug("unregister, rdb delete error")
	}
}
