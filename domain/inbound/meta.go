package inbound

import (
	"github.com/ppzxc/chattools/common"
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

func (m Meta) GetMetaType() common.Meta {
	if m.Topic != nil {
		return common.MetaTopic
	} else if m.User != nil {
		return common.MetaUser
	} else if m.Message != nil {
		return common.MetaMessage
	} else if m.Notify != nil {
		return common.MetaNotify
	} else if m.Profile != nil {
		return common.MetaProfile
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
	TopicId    int64         `json:"topic_id,omitempty"`
	SequenceId int64         `json:"sequence_id,omitempty"`
	Paging     *model.Paging `json:"paging,omitempty"`
}

type RequestNotify struct {
	Paging *model.Paging `json:"paging,omitempty"`
}

type RequestProfile struct {
	UserId      int64  `json:"user_id,omitempty"`
	FileId      int64  `json:"file_id,omitempty"`
	Description string `json:"description,omitempty" validate:"required"`
}
