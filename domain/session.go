package domain

import "github.com/json-iterator/go"

type SessionAdapter interface {
	GetSessionId() string

	IsLogin() bool
	Login(int64, string)
	Logout()

	GetUserId() int64
	GetBrowserId() string

	ToExternal() []byte
}

func NewSessionFromExternal(payload interface{}) SessionAdapter {
	var ls session
	_ = jsoniter.Unmarshal(payload.([]byte), &ls)
	return &ls
}

func NewSession(sessionId string) SessionAdapter {
	return &session{
		SessionId:  sessionId,
		LoginState: false,
		UserId:     0,
		BrowserId:  "",
	}
}

type session struct {
	LoginState bool   `json:"login_state"`
	SessionId  string `json:"session_id"`
	UserId     int64  `json:"user_id"`
	BrowserId  string `json:"browser_id"`
}

func (s session) ToExternal() []byte {
	marshal, _ := jsoniter.Marshal(s)
	return marshal
}

func (s session) GetSessionId() string {
	return s.SessionId
}

func (s session) IsLogin() bool {
	return s.LoginState
}

func (s *session) Login(userId int64, browserId string) {
	s.LoginState = true
	s.UserId = userId
	s.BrowserId = browserId
}

func (s *session) Logout() {
	s.LoginState = false
	s.UserId = 0
	s.BrowserId = ""
}

func (s session) GetUserId() int64 {
	return s.UserId
}

func (s session) GetBrowserId() string {
	return s.BrowserId
}
