package outbound

import "time"

type Notify struct {
	UUID    string           `json:"uuid,omitempty"`
	Mention *ResponseMention `json:"mention,omitempty"`
	Reply   *ResponseReply   `json:"reply,omitempty"`
}

type ResponseMention struct {
	Create  *ResponseNotifyCreate `json:"create,omitempty"`
	Receive *ResponseNotifyCrud   `json:"receive,omitempty"`
	Read    *ResponseNotifyCrud   `json:"read,omitempty"`
	Delete  *ResponseNotifyCrud   `json:"delete,omitempty"`
}

type ResponseReply struct {
	Create  *ResponseNotifyCreate `json:"create,omitempty"`
	Receive *ResponseNotifyCrud   `json:"receive,omitempty"`
	Read    *ResponseNotifyCrud   `json:"read,omitempty"`
	Delete  *ResponseNotifyCrud   `json:"delete,omitempty"`
}

type ResponseNotifyCrud struct {
	NotifyId int64 `json:"notify_id,omitempty"`
}

type ResponseNotifyCreate struct {
	Notification *Notification `json:"notification,omitempty"`
}

type Notification struct {
	Id            int64       `json:"id,omitempty"`
	Type          string      `json:"type,omitempty"`
	RequestUserId int64       `json:"request_user_id,omitempty"`
	ReceiveUserId int64       `json:"receive_user_id,omitempty"`
	TopicId       int64       `json:"topic_id,omitempty"`
	SequenceId    int64       `json:"sequence_id,omitempty"`
	IsReceived    bool        `json:"is_received"`
	IsRead        bool        `json:"is_read"`
	Custom        interface{} `json:"custom,omitempty"`
	CreatedAt     *time.Time  `json:"created_at,omitempty"`
	UpdatedAt     *time.Time  `json:"updated_at,omitempty"`
}
