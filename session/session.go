package session

import "github.com/json-iterator/go"

type Session interface {
	GetSessionId() string

	IsLogin() bool
	Login(int64, int64)
	Logout()

	GetUserId() int64
	GetDeviceId() int64

	ToExternal() []byte
}

func NewSessionFromExternal(payload interface{}) Session {
	var ls localSession
	_ = jsoniter.Unmarshal(payload.([]byte), &ls)
	return &ls
}

func NewSession(sessionId string) Session {
	return &localSession{
		SessionId:  sessionId,
		LoginState: false,
		UserId:     0,
		DeviceId:   0,
	}
}

type localSession struct {
	LoginState bool   `json:"login_state"`
	SessionId  string `json:"session_id"`
	UserId     int64  `json:"user_id"`
	DeviceId   int64  `json:"device_id"`
}

func (s localSession) ToExternal() []byte {
	marshal, _ := jsoniter.Marshal(s)
	return marshal
}

func (s localSession) GetSessionId() string {
	return s.SessionId
}

func (s localSession) IsLogin() bool {
	return s.LoginState
}

func (s *localSession) Login(userId int64, deviceId int64) {
	s.LoginState = true
	s.UserId = userId
	s.DeviceId = deviceId
}

func (s *localSession) Logout() {
	s.LoginState = false
	s.UserId = 0
	s.DeviceId = 0
}

func (s localSession) GetUserId() int64 {
	return s.UserId
}

func (s localSession) GetDeviceId() int64 {
	return s.DeviceId
}
