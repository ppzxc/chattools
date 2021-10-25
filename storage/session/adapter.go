package session

import (
	"context"
	"errors"
	"github.com/kataras/iris/v12/websocket"
	"github.com/ppzxc/chattools/domain"
)

var (
	ErrContainsSessionStore       = errors.New("session register failed, contains session id")
	ErrNotRegister                = errors.New("session is not register store")
	ErrAlreadyRegister            = errors.New("session is already register")
	ErrAlreadyLogin               = errors.New("session is already login")
	ErrNoLoginState               = errors.New("session is not login state")
	ErrUserSessionAlreadySessions = errors.New("already session in user session store")
)

type Adapter interface {
	Login(sessionId string, userId int64, deviceId string) error
	Logout(sessionId string) error
	GetSession(sessionId string) (domain.SessionAdapter, bool)
	GetSessionByUserId(userId int64) (map[string]domain.SessionAdapter, error)
	Register(session domain.SessionAdapter) error
	Unregister(sessionId string)

	Subscribe(ctx context.Context, key string, conn *websocket.Conn) error
	Publish(ctx context.Context, key string, message interface{}) error
}
