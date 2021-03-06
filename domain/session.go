package domain

import (
	"github.com/json-iterator/go"
	"strconv"
)

type SessionAdapter interface {
	GetSessionId() string

	IsLogin() bool
	Login(int64, string, string)
	Logout()

	GetUserId() int64
	GetUserName() string
	GetBrowserId() string

	ToMap() map[string]interface{}
}

func FromHash(payload interface{}) SessionAdapter {
	//fmt.Println("from", reflect.TypeOf(payload), string(payload.([]byte)))
	var ls session
	_ = jsoniter.Unmarshal(payload.([]byte), &ls)
	return &ls
}

func (s session) ToMap() map[string]interface{} {
	sess := make(map[string]interface{})
	sess["session_id"] = s.SessionId

	if s.LoginState {
		sess["login_state"] = s.LoginState
	}

	if s.UserId > 0 {
		sess["user_id"] = s.UserId
	}

	if len(s.UserName) > 0 {
		sess["user_name"] = s.UserName
	}

	if len(s.BrowserId) > 0 {
		sess["browser_id"] = s.BrowserId
	}
	return sess
}

func FromMap(payload map[string]string) (SessionAdapter, error) {
	sess := session{}
	value, loaded := payload["login_state"]
	if loaded {
		isLogin, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		sess.LoginState = isLogin
	}

	value, loaded = payload["user_id"]
	if loaded {
		userId, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		sess.UserId = int64(userId)
	}

	value, loaded = payload["user_name"]
	if loaded {
		sess.UserName = value
	}

	value, loaded = payload["session_id"]
	if loaded {
		sess.SessionId = value
	}

	value, loaded = payload["browser_id"]
	if loaded {
		sess.BrowserId = value
	}
	return &sess, nil
}

func NewSession(sessionId string) SessionAdapter {
	return &session{
		SessionId:  sessionId,
		LoginState: false,
		UserId:     0,
		UserName:   "",
		BrowserId:  "",
	}
}

type session struct {
	LoginState bool   `json:"login_state"`
	SessionId  string `json:"session_id"`
	UserId     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	BrowserId  string `json:"browser_id"`
}

func (s session) GetSessionId() string {
	return s.SessionId
}

func (s session) IsLogin() bool {
	return s.LoginState
}

func (s *session) Login(userId int64, userName string, browserId string) {
	s.LoginState = true
	s.UserId = userId
	s.UserName = userName
	s.BrowserId = browserId
}

func (s *session) Logout() {
	s.LoginState = false
	s.UserId = 0
	s.UserName = ""
	s.BrowserId = ""
}

func (s session) GetUserId() int64 {
	return s.UserId
}

func (s session) GetUserName() string {
	return s.UserName
}

func (s session) GetBrowserId() string {
	return s.BrowserId
}
