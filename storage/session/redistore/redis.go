package redistore

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ppzxc/chattools/domain"
	"github.com/ppzxc/chattools/storage/cache"
	"github.com/ppzxc/chattools/storage/session"
	"github.com/ppzxc/chattools/types"
	"github.com/ppzxc/chattools/utils"
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

func (r *redisSessionStore) Subscribe(ctx context.Context, key ...string) (*redis.PubSub, error) {
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("subscribe called")
	return r.rdb.Subscribe(ctx, key...)
}

func (r *redisSessionStore) Publish(ctx context.Context, key string, message interface{}) error {
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("publish called")
	return r.rdb.Publish(ctx, key, message)
}

func (r *redisSessionStore) Login(ctx context.Context, sessionId string, userId int64, userName string, browserId string) error {
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("login called")
	if err := r.rdb.Exists(ctx, sessionId); err != nil {
		return err
	}

	sess := domain.NewSession(sessionId)
	sess.Login(userId, userName, browserId)

	err := r.rdb.HSet(ctx, sessionId, sess.ToMap())
	if err != nil {
		return err
	}

	if err = r.rdb.Exists(ctx, getUserKey(userId)); err != nil {
		if err != types.ErrNoExistsKeys {
			return err
		}
	}

	return r.rdb.HSet(ctx, getUserKey(userId), sessionId, browserId)
}

func (r *redisSessionStore) Logout(ctx context.Context, sessionId string) error {
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("logout called")
	get, err := r.rdb.HGetAll(ctx, sessionId)
	if err != nil {
		return err
	}

	sess, err := domain.FromMap(get)
	if err != nil {
		return err
	}

	if sess.IsLogin() {
		if err := r.rdb.HExists(ctx, getUserKey(sess.GetUserId()), sessionId); err == nil {
			_ = r.rdb.HDel(ctx, getUserKey(sess.GetUserId()), sessionId)
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
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("ExistSession called")
	err := r.rdb.Exists(ctx, sessionId)
	if err != nil {
		return false
	}
	return true
}

func (r *redisSessionStore) GetSession(ctx context.Context, sessionId string) (domain.SessionAdapter, bool) {
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("GetSession called")
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
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("GetSessionByUserId called")
	maps, err := r.rdb.HGetAll(ctx, getUserKey(userId))
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
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("Register called")
	err := r.rdb.Exists(ctx, registerSession.GetSessionId())
	if err != nil && err != types.ErrNoExistsKeys {
		return err
	}

	return r.rdb.HSet(ctx, registerSession.GetSessionId(), registerSession.ToMap())
}

func (r *redisSessionStore) Unregister(ctx context.Context, sessionId string) {
	logrus.WithFields(utils.ContextValueExtractor(ctx, logrus.Fields{})).Debug("Unregister called")
	get, err := r.rdb.HGetAll(ctx, sessionId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"session.id": sessionId,
		}).WithError(err).Error("unregister session get failed")
		return
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
		logrus.WithFields(logrus.Fields{
			"session.id": sessionId,
		}).WithError(err).Error("unregister, fromMap transform fail")
		return
	}

	if !sess.IsLogin() {
		logrus.WithFields(logrus.Fields{
			"session.id": sessionId,
		}).WithError(err).Error("session is not login")
		return
	}

	all, err := r.rdb.HGetAll(ctx, getUserKey(sess.GetUserId()))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"session.id": sessionId,
		}).WithError(err).Error("unregister user sessions get failed")
		return
	}

	if browserId, loaded := all[sessionId]; loaded {
		if len(all) <= 1 {
			if err := r.rdb.Del(ctx, getUserKey(sess.GetUserId())); err != nil {
				logrus.WithFields(logrus.Fields{
					"session.id": sessionId,
					"browser.id": browserId,
				}).WithError(err).Error("unregister user sessions get failed")
				return
			}
		} else {
			if err := r.rdb.HDel(ctx, getUserKey(sess.GetUserId()), sess.GetSessionId()); err != nil {
				logrus.WithFields(logrus.Fields{
					"session.id": sessionId,
					"browser.id": browserId,
				}).WithError(err).Error("unregister user sessions get failed")
				return
			}
		}
	}
}
