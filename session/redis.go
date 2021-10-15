package session

import (
	"github.com/ppzxc/chattools/cache"
	"github.com/sirupsen/logrus"
)

func NewRedisSessionStore(address string, password string, db int, isReset bool) Adapter {
	return &redisSessionStore{
		rdb: cache.NewRedisCache(address, password, db, isReset),
	}
}

type redisSessionStore struct {
	rdb cache.Adapter
}

func (r redisSessionStore) Login(sessionId string, userId int64, deviceId int64) error {
	get, err := r.rdb.Get(sessionId)
	if err != nil || get == nil {
		return ErrNotRegister
	} else {
		sess := NewSessionFromExternal(get)
		sess.Login(userId, deviceId)
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
		return ErrNotRegister
	} else {
		sess := NewSessionFromExternal(get)
		sess.Logout()
		err := r.rdb.Set(sessionId, sess.ToExternal())
		if err != nil {
			return err
		}
		return nil
	}
}

func (r redisSessionStore) GetSession(sessionId string) (Session, bool) {
	get, err := r.rdb.Get(sessionId)
	if err != nil || get == nil {
		return nil, false
	} else {
		return NewSessionFromExternal(get), true
	}
}

//TODO get user sessions in multidevce
func (r redisSessionStore) GetSessions(userId int64) (map[string]Session, bool) {
	panic("implement me")
}

func (r redisSessionStore) Register(session Session) error {
	get, err := r.rdb.Get(session.GetSessionId())
	if err == nil && get != nil {
		return ErrAlreadyRegister
	} else {
		err := r.rdb.Set(session.GetSessionId(), session.ToExternal())
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
