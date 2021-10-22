package localstore

//
//import (
//	"context"
//	"github.com/go-redis/redis/v8"
//	"github.com/ppzxc/chattools/domain"
//	session2 "github.com/ppzxc/chattools/storage/session"
//	"sync"
//)
//
//func NewLocalSessionStore() session2.Adapter {
//	ss := &sessionStore{
//		sessions:     make(map[string]domain.SessionAdapter),
//		userSessions: make(map[int64]map[string]domain.SessionAdapter),
//	}
//	return ss
//}
//
//type sessionStore struct {
//	sync.Mutex
//	sessions     map[string]domain.SessionAdapter
//	userSessions map[int64]map[string]domain.SessionAdapter
//}
//
//func (s *sessionStore) Subscribe(ctx context.Context, key string) (*redis.PubSub, error) {
//	panic("implement me")
//}
//
//func (s *sessionStore) Publish(ctx context.Context, key string, message interface{}) error {
//	panic("implement me")
//}
//
//func (s *sessionStore) Login(sessionId string, userId int64, browserId string) error {
//	s.Lock()
//	defer s.Unlock()
//
//	cacheSession, loaded := s.sessions[sessionId]
//	if !loaded {
//		return session2.ErrNotRegister
//	}
//
//	if cacheSession.IsLogin() {
//		return session2.ErrAlreadyLogin
//	}
//
//	cacheSession.Login(userId, browserId)
//
//	if sess, l := s.userSessions[userId]; l {
//		_, loaded = sess[sessionId]
//		if loaded {
//			return session2.ErrUserSessionAlreadySessions
//		} else {
//			sess[sessionId] = cacheSession
//			return nil
//		}
//	} else {
//		input := make(map[string]domain.SessionAdapter)
//		input[sessionId] = cacheSession
//		s.userSessions[userId] = input
//		return nil
//	}
//}
//
//func (s *sessionStore) Logout(sessionId string) error {
//	s.Lock()
//	defer s.Unlock()
//
//	cacheSession, loaded := s.sessions[sessionId]
//	if !loaded {
//		return session2.ErrNotRegister
//	} else {
//		if !cacheSession.IsLogin() {
//			return session2.ErrNoLoginState
//		}
//
//		if us, loaded := s.userSessions[cacheSession.GetUserId()]; loaded {
//			_, loaded := us[sessionId]
//			if loaded {
//				delete(us, sessionId)
//			}
//		}
//
//		cacheSession.Logout()
//		return nil
//	}
//}
//
//func (s *sessionStore) GetSession(sessionId string) (domain.SessionAdapter, bool) {
//	s.Lock()
//	defer s.Unlock()
//	cacheSession, loaded := s.sessions[sessionId]
//	return cacheSession, loaded
//}
//
//func (s *sessionStore) GetSessions(userId int64) (map[string]domain.SessionAdapter, bool) {
//	s.Lock()
//	defer s.Unlock()
//	cacheSession, loaded := s.userSessions[userId]
//	return cacheSession, loaded
//}
//
//func (s *sessionStore) Register(registerSession domain.SessionAdapter) error {
//	s.Lock()
//	defer s.Unlock()
//
//	_, loaded := s.sessions[registerSession.GetSessionId()]
//	if loaded {
//		return session2.ErrContainsSessionStore
//	} else {
//		s.sessions[registerSession.GetSessionId()] = registerSession
//		return nil
//	}
//}
//
//func (s *sessionStore) Unregister(sessionId string) {
//	s.Lock()
//	defer s.Unlock()
//
//	_, loaded := s.sessions[sessionId]
//	if loaded {
//		delete(s.sessions, sessionId)
//	}
//}
