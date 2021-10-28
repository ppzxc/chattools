package inbound

import "github.com/ppzxc/chattools/common/global"

type Message struct {
	UUID string       `json:"uuid,omitempty" validate:"required,uuid4"`
	Send *RequestSend `json:"create,omitempty"`
	Ack  *RequestAck  `json:"leave,omitempty"`
	Read *RequestRead `json:"invite,omitempty"`
	File *RequestFile `json:"join,omitempty"`
}

func (m Message) GetMsgType() global.Msg {
	if m.Send != nil {
		return global.MsgSend
	} else if m.Ack != nil {
		return global.MsgAck
	} else if m.Read != nil {
		return global.MsgRead
	} else if m.File != nil {
		return global.MsgFile
	} else {
		return ""
	}
}

type RequestSend struct {
	TopicId int64  `json:"topic_id,omitempty" validate:"required,min=1"`
	Message string `json:"message,omitempty" validate:"required,min=1,max=2000"`
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
