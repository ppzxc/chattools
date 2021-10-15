package session

import (
	"sync"
)

func NewLocalSessionStore() Adapter {
	ss := &sessionStore{
		sessions:     make(map[string]Session),
		userSessions: make(map[int64]map[string]Session),
	}
	return ss
}

type sessionStore struct {
	sync.Mutex
	sessions     map[string]Session
	userSessions map[int64]map[string]Session
}

func (s *sessionStore) Login(sessionId string, userId int64, deviceId int64) error {
	s.Lock()
	defer s.Unlock()

	session, loaded := s.sessions[sessionId]
	if !loaded {
		return ErrNotRegister
	}

	if session.IsLogin() {
		return ErrAlreadyLogin
	}

	session.Login(userId, deviceId)

	if sess, l := s.userSessions[userId]; l {
		_, loaded = sess[sessionId]
		if loaded {
			return ErrUserSessionAlreadySessions
		} else {
			sess[sessionId] = session
			return nil
		}
	} else {
		input := make(map[string]Session)
		input[sessionId] = session
		s.userSessions[userId] = input
		return nil
	}
}

func (s *sessionStore) Logout(sessionId string) error {
	s.Lock()
	defer s.Unlock()

	session, loaded := s.sessions[sessionId]
	if !loaded {
		return ErrNotRegister
	} else {
		if !session.IsLogin() {
			return ErrNoLoginState
		}

		us, loaded := s.userSessions[session.GetUserId()]
		if loaded {
			_, loaded := us[sessionId]
			if loaded {
				delete(us, sessionId)
			}
		}
		session.Logout()
		return nil
	}
}

func (s *sessionStore) GetSession(sessionId string) (Session, bool) {
	s.Lock()
	defer s.Unlock()

	session, loaded := s.sessions[sessionId]
	return session, loaded
}

func (s *sessionStore) GetSessions(userId int64) (map[string]Session, bool) {
	s.Lock()
	defer s.Unlock()

	session, loaded := s.userSessions[userId]
	return session, loaded
}

func (s *sessionStore) Register(session Session) error {
	s.Lock()
	defer s.Unlock()

	_, loaded := s.sessions[session.GetSessionId()]
	if loaded {
		return ErrContainsSessionStore
	} else {
		s.sessions[session.GetSessionId()] = session
		return nil
	}
}

func (s *sessionStore) Unregister(sessionId string) {
	s.Lock()
	defer s.Unlock()

	_, loaded := s.sessions[sessionId]
	if loaded {
		delete(s.sessions, sessionId)
	}
}
