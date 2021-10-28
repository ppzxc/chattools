package inbound

import (
	"github.com/ppzxc/chattools/common/global"
	"github.com/ppzxc/chattools/storage/database/model"
)

type Meta struct {
	UUID    string          `json:"uuid,omitempty" validate:"required,uuid4"`
	Topic   *RequestTopic   `json:"topic,omitempty"`
	User    *RequestUser    `json:"user,omitempty"`
	Message *RequestMessage `json:"message,omitempty"`
	Notify  *RequestNotify  `json:"notify,omitempty"`
	Profile *RequestProfile `json:"profile,omitempty"`
}

func (m Meta) GetMetaType() global.Meta {
	if m.Topic != nil {
		return global.MetaTopic
	} else if m.User != nil {
		return global.MetaUser
	} else if m.Message != nil {
		return global.MetaMessage
	} else if m.Notify != nil {
		return global.MetaNotify
	} else if m.Profile != nil {
		return global.MetaProfile
	} else {
		return ""
	}
}

type RequestTopic struct {
	Mine   bool          `json:"mine" validate:"required"`
	Paging *model.Paging `json:"paging,omitempty"`
}

type RequestUser struct {
	TopicId int64         `json:"topic_id,omitempty"`
	Paging  *model.Paging `json:"paging,omitempty"`
}

type RequestMessage struct {
}

type RequestNotify struct {
}

type RequestProfile struct {
}
