package session

import "github.com/json-iterator/go"

type Session interface {
	GetSessionId() string

	IsLogin() bool
	Login(int64, string)
	Logout()

	GetUserId() int64
	GetBrowserId() string

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
		BrowserId:  "",
	}
}

type localSession struct {
	LoginState bool   `json:"login_state"`
	SessionId  string `json:"session_id"`
	UserId     int64  `json:"user_id"`
	BrowserId  string `json:"browser_id"`
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

func (s *localSession) Login(userId int64, browserId string) {
	s.LoginState = true
	s.UserId = userId
	s.BrowserId = browserId
}

func (s *localSession) Logout() {
	s.LoginState = false
	s.UserId = 0
	s.BrowserId = ""
}

func (s localSession) GetUserId() int64 {
	return s.UserId
}

func (s localSession) GetBrowserId() string {
	return s.BrowserId
}
