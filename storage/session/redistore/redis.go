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
)

func getUserKey(userId int64) string {
	return fmt.Sprintf("USER_%v", userId)
}

func NewRedisSessionStore(adapter cache.Adapter) session.Adapter {
	return &redisSessionStore{
		rdb: adapter,
	}
}

type redisSessionStore struct {
	rdb cache.Adapter
}

func (r *redisSessionStore) Subscribe(ctx context.Context, key string) (*redis.PubSub, error) {
	return r.rdb.Subscribe(ctx, key)
}

func (r *redisSessionStore) Publish(ctx context.Context, key string, message interface{}) error {
	return r.rdb.Publish(ctx, key, message)
}

func (r *redisSessionStore) Login(sessionId string, userId int64, browserId string) error {
	if err := r.rdb.Exists(sessionId); err != nil {
		return err
	}

	sess := domain.NewSession(sessionId)
	sess.Login(userId, browserId)

	err := r.rdb.HSet(sessionId, sess.ToMap())
	if err != nil {
		return err
	}

	if err = r.rdb.Exists(getUserKey(userId)); err != nil {
		if err != types.ErrNoExistsKeys {
			return err
		}
	} // fallthrough

	return r.rdb.HSet(getUserKey(userId), sessionId, browserId)
}

func (r *redisSessionStore) Logout(sessionId string) error {
	get, err := r.rdb.HGetAll(sessionId)
	if err != nil {
		return err
	}

	sess, err := domain.FromMap(get)
	if err != nil {
		return err
	}

	if sess.IsLogin() {
		if err := r.rdb.HExists(getUserKey(sess.GetUserId()), sessionId); err == nil {
			_ = r.rdb.HDel(getUserKey(sess.GetUserId()), sessionId)
		}

		sess.Logout()
		err := r.rdb.HDel(sessionId, "login_state", "user_id", "browser_id")
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *redisSessionStore) GetSession(sessionId string) (domain.SessionAdapter, bool) {
	get, err := r.rdb.HGetAll(sessionId)
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

func (r *redisSessionStore) GetSessionByUserId(userId int64) (map[string]domain.SessionAdapter, error) {
	maps, err := r.rdb.HGetAll(getUserKey(userId))
	if err != nil {
		return nil, err
	}

	s := make(map[string]domain.SessionAdapter)
	for key := range maps {
		all, err := r.rdb.HGetAll(key)
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

func (r *redisSessionStore) Register(registerSession domain.SessionAdapter) error {
	if err := r.rdb.Exists(registerSession.GetSessionId()); err != nil {
		if err != types.ErrNoExistsKeys {
			return err
		}
	}

	return r.rdb.HSet(registerSession.GetSessionId(), registerSession.ToMap())
}

func (r *redisSessionStore) Unregister(sessionId string) {
	get, err := r.rdb.HGetAll(sessionId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"session.id": sessionId,
		}).WithError(err).Error("unregister session get failed")
		return
	}

	defer func() {
		err := r.rdb.Del(sessionId)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"session.id": sessionId,
			}).WithError(err).Error("unregister, rdb delete error")
		}
	}()

	sess, err := domain.FromMap(get)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"session.id": sessionId,
		}).WithError(err).Error("unregister, fromMap transform fail")
		return
	}

	if sess.IsLogin() {
		all, err := r.rdb.HGetAll(getUserKey(sess.GetUserId()))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"session.id": sessionId,
			}).WithError(err).Error("unregister user sessions get failed")
			return
		}

		if _, loaded := all[sessionId]; loaded {
			if len(all) <= 1 {
				if err := r.rdb.Del(getUserKey(sess.GetUserId())); err != nil {
					logrus.WithFields(logrus.Fields{
						"session.id": sessionId,
					}).WithError(err).Error("unregister user sessions get failed")
					return
				}
			} else {
				err := r.rdb.HDel(getUserKey(sess.GetUserId()), sess.GetSessionId())
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"session.id": sessionId,
					}).WithError(err).Error("unregister user sessions get failed")
					return
				}
			}
		}
	}
}
