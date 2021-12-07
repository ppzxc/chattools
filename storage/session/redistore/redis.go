package redistore

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ppzxc/chattools/domain"
	"github.com/ppzxc/chattools/storage/cache"
	"github.com/ppzxc/chattools/storage/session"
	"github.com/ppzxc/chattools/types"
	"github.com/sirupsen/logrus"
	"sync"
)

func NewRedisSessionStore(adapter cache.Adapter) session.Adapter {
	return &redisSessionStore{
		rdb: adapter,
		mtx: sync.Mutex{},
		pss: make(map[string]*redis.PubSub),
	}
}

type redisSessionStore struct {
	rdb cache.Adapter
	mtx sync.Mutex
	pss map[string]*redis.PubSub
}

func (r *redisSessionStore) UnsubscribeUser(ctx context.Context, sessionId string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if ps, ok := r.pss[sessionId]; ok {
		err := ps.Unsubscribe(ctx)
		if err != nil {
			return err
		}
		err = ps.Close()
		if err != nil {
			return err
		}
		delete(r.pss, sessionId)
		return nil
	}
	return nil
}

func (r *redisSessionStore) UnsubscribeTopic(ctx context.Context, sessionId string, topicId int64) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if ps, ok := r.pss[sessionId]; ok {
		err := ps.Unsubscribe(ctx, r.GetTopicKey(topicId))
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (r *redisSessionStore) PubSubNumSub(ctx context.Context, key ...string) (map[string]int64, error) {
	return r.rdb.PubSubNumSub(ctx, key...)
}

func (r *redisSessionStore) SubscribeTopic(ctx context.Context, sessionId string, userId int64, topicId int64) (<-chan *redis.Message, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if ps, ok := r.pss[sessionId]; ok {
		err := ps.Subscribe(ctx, r.GetTopicKey(topicId))
		if err != nil {
			return nil, err
		}
		return ps.Channel(), nil
	} else {
		subscribe, err := r.rdb.Subscribe(ctx, r.GetUserKey(userId))
		if err != nil {
			return nil, err
		}
		err = subscribe.Subscribe(ctx, r.GetTopicKey(topicId))
		if err != nil {
			return nil, err
		}
		r.pss[sessionId] = subscribe
		return subscribe.Channel(), nil
	}
}

func (r *redisSessionStore) GetTopicKey(topicId int64) string {
	return fmt.Sprintf("TOPIC.%v", topicId)
}

func (r *redisSessionStore) GetUserKey(userId int64) string {
	return fmt.Sprintf("USER.%v", userId)
}

func (r *redisSessionStore) SubscribeUser(ctx context.Context, sessionId string, userId int64) (<-chan *redis.Message, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if ps, ok := r.pss[sessionId]; ok {
		err := ps.Subscribe(ctx, r.GetUserKey(userId))
		if err != nil {
			return nil, err
		}
		return ps.Channel(), nil
	} else {
		subscribe, err := r.rdb.Subscribe(ctx, r.GetUserKey(userId))
		if err != nil {
			return nil, err
		}
		r.pss[sessionId] = subscribe
		return subscribe.Channel(), nil
	}
}

func (r *redisSessionStore) Publish(ctx context.Context, key string, message interface{}) error {
	return r.rdb.Publish(ctx, key, message)
}

func (r *redisSessionStore) Login(ctx context.Context, sessionId string, userId int64, userName string, browserId string) error {
	if err := r.rdb.Exists(ctx, sessionId); err != nil {
		return err
	}

	sess := domain.NewSession(sessionId)
	sess.Login(userId, userName, browserId)

	err := r.rdb.HSet(ctx, sessionId, sess.ToMap())
	if err != nil {
		return err
	}

	if err = r.rdb.Exists(ctx, r.GetUserKey(userId)); err != nil {
		if err != types.ErrNoExistsKeys {
			return err
		}
	}

	return r.rdb.HSet(ctx, r.GetUserKey(userId), sessionId, browserId)
}

func (r *redisSessionStore) Logout(ctx context.Context, sessionId string) error {
	get, err := r.rdb.HGetAll(ctx, sessionId)
	if err != nil {
		return err
	}

	sess, err := domain.FromMap(get)
	if err != nil {
		return err
	}

	if sess.IsLogin() {
		if err := r.rdb.HExists(ctx, r.GetUserKey(sess.GetUserId()), sessionId); err == nil {
			_ = r.rdb.HDel(ctx, r.GetUserKey(sess.GetUserId()), sessionId)
		}

		sess.Logout()
		err := r.rdb.HDel(ctx, sessionId, "login_state", "user_id", "browser_id", "user_name")
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *redisSessionStore) ExistSession(ctx context.Context, sessionId string) bool {
	err := r.rdb.Exists(ctx, sessionId)
	if err != nil {
		return false
	}
	return true
}

func (r *redisSessionStore) GetSession(ctx context.Context, sessionId string) (domain.SessionAdapter, bool) {
	get, err := r.rdb.HGetAll(ctx, sessionId)
	if err != nil || get == nil {
		return nil, false
	} else {
		fromMap, err := domain.FromMap(get)
		if err != nil {
			return nil, false
		}
		return fromMap, true
	}
}

func (r *redisSessionStore) GetSessionByUserId(ctx context.Context, userId int64) (map[string]domain.SessionAdapter, error) {
	maps, err := r.rdb.HGetAll(ctx, r.GetUserKey(userId))
	if err != nil {
		return nil, err
	}

	s := make(map[string]domain.SessionAdapter)
	for key := range maps {
		all, err := r.rdb.HGetAll(ctx, key)
		if err != nil {
			return nil, err
		}

		fromMap, err := domain.FromMap(all)
		if err != nil {
			return nil, err
		}
		s[key] = fromMap
	}

	return s, nil
}

func (r *redisSessionStore) Register(ctx context.Context, registerSession domain.SessionAdapter) error {
	err := r.rdb.Exists(ctx, registerSession.GetSessionId())
	if err != nil && err != types.ErrNoExistsKeys {
		return err
	}
	return r.rdb.HSet(ctx, registerSession.GetSessionId(), registerSession.ToMap())
}

func (r *redisSessionStore) Unregister(ctx context.Context, sessionId string) error {
	get, err := r.rdb.HGetAll(ctx, sessionId)
	if err != nil {
		return err
	}

	defer func() {
		if err := r.rdb.Del(ctx, sessionId); err != nil {
			logrus.WithFields(logrus.Fields{
				"session.id": sessionId,
			}).WithError(err).Error("unregister, rdb delete error")
		} else {
			logrus.WithFields(logrus.Fields{
				"session.id": sessionId,
			}).Info("unregister done")
		}
	}()

	sess, err := domain.FromMap(get)
	if err != nil {
		return err
	}

	if !sess.IsLogin() {
		return nil
	}

	all, err := r.rdb.HGetAll(ctx, r.GetUserKey(sess.GetUserId()))
	if err != nil {
		return err
	}

	if _, loaded := all[sessionId]; loaded {
		if len(all) <= 1 {
			if err := r.rdb.Del(ctx, r.GetUserKey(sess.GetUserId())); err != nil {
				return err
			}
		} else {
			if err := r.rdb.HDel(ctx, r.GetUserKey(sess.GetUserId()), sess.GetSessionId()); err != nil {
				return err
			}
		}
	}
	return nil
}
