package session

import "errors"

var (
	ErrContainsSessionStore       = errors.New("session register failed, contains session id")
	ErrNotRegister                = errors.New("session is not register store")
	ErrAlreadyRegister            = errors.New("session is already register")
	ErrAlreadyLogin               = errors.New("session is already login")
	ErrNoLoginState               = errors.New("session is not login state")
	ErrUserSessionAlreadySessions = errors.New("already session in user session store")
)

type Adapter interface {
	Login(sessionId string, userId int64, deviceId int64) error
	Logout(sessionId string) error
	GetSession(sessionId string) (Session, bool)
	GetSessions(userId int64) (map[string]Session, bool)
	Register(session Session) error
	Unregister(sessionId string)
}