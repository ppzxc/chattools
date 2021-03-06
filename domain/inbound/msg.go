package inbound

import (
	"github.com/ppzxc/chattools/common"
)

type Message struct {
	UUID string       `json:"uuid,omitempty" validate:"required,uuid4"`
	Send *RequestSend `json:"send,omitempty"`
	Ack  *RequestAck  `json:"ack,omitempty"`
	Read *RequestRead `json:"read,omitempty"`
	File *RequestFile `json:"file,omitempty"`
}

func (m Message) GetMsgType() common.Msg {
	if m.Send != nil {
		return common.MsgSend
	} else if m.Ack != nil {
		return common.MsgAck
	} else if m.Read != nil {
		return common.MsgRead
	} else if m.File != nil {
		return common.MsgFile
	} else {
		return ""
	}
}

type RequestSend struct {
	TopicId int64       `json:"topic_id,omitempty" validate:"required,min=1"`
	Message string      `json:"message,omitempty" validate:"required,min=1,max=2000"`
	Custom  interface{} `json:"custom,omitempty"`
}

type RequestAck struct {
	TopicId    int64 `json:"topic_id,omitempty" validate:"required,min=1"`
	SequenceId int64 `json:"sequence_id,omitempty" validate:"required,min=1"`
}

type RequestRead struct {
	TopicId    int64 `json:"topic_id,omitempty" validate:"required,min=1"`
	SequenceId int64 `json:"sequence_id,omitempty" validate:"required,min=1"`
}

type RequestFile struct {
	TopicId int64 `json:"topic_id,omitempty" validate:"required,min=1"`
	FileId  int64 `json:"file_id,omitempty" validate:"required,min=1"`
}
