package outbound

import "time"

type Control struct {
	UUID   string        `json:"uuid,omitempty"`
	Create *ResponseCtrl `json:"create,omitempty"`
	Leave  *ResponseCtrl `json:"leave,omitempty"`
	Invite *ResponseCtrl `json:"invite,omitempty"`
	Join   *ResponseCtrl `json:"join,omitempty"`
}

type ResponseCtrl struct {
	Topic   *Topic `json:"topic,omitempty"`
	Name    string `json:"name,omitempty"`
	TopicId int64  `json:"topic_id,omitempty"`
	UserId  int64  `json:"user_id,omitempty"`
}

type Topic struct {
	Id        int64      `json:"id,omitempty"`
	State     string     `json:"state,omitempty"`
	StatedAt  *time.Time `json:"stated_at,omitempty"`
	Name      string     `json:"name,omitempty"`
	Owner     int64      `json:"owner,omitempty"`
	Private   bool       `json:"private,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Message []*Message `json:"messages"`
	Users   []*User    `json:"users"`
}

type Message struct {
	Id           int64       `json:"id,omitempty"`
	MessageType  string      `json:"message_type,omitempty"`
	ClientUUID   string      `json:"client_uuid,omitempty"`
	FromUserId   int64       `json:"from_user_id,omitempty"`
	FromUserName string      `json:"from_user_name,omitempty"`
	TopicId      int64       `json:"topic_id,omitempty"`
	FileId       int64       `json:"file_id,omitempty"`
	SequenceId   int64       `json:"sequence_id,omitempty"`
	Content      string      `json:"content,omitempty"`
	Custom       interface{} `json:"custom,omitempty"`
	CreatedAt    *time.Time  `json:"created_at,omitempty"`
	UpdatedAt    *time.Time  `json:"updated_at,omitempty"`
	DeletedAt    *time.Time  `json:"deleted_at,omitempty"`
}
