package inbound

import "github.com/ppzxc/chattools/common/global"

type Notify struct {
	UUID    string          `json:"uuid,omitempty" validate:"required,uuid4"`
	Mention *RequestMention `json:"mention,omitempty"`
	Reply   *RequestReply   `json:"reply,omitempty"`
}

func (n Notify) GetNotifyType() global.Notify {
	if n.Mention != nil {
		return global.NotifyMention
	} else if n.Reply != nil {
		return global.NotifyReply
	} else {
		return ""
	}
}

type RequestMention struct {
	Create  *MentionCreate `json:"create,omitempty"`
	Receive *Crud          `json:"receive,omitempty"`
	Read    *Crud          `json:"read,omitempty"`
	Delete  *Crud          `json:"delete,omitempty"`
}

func (r RequestMention) GetNotifyType() global.NotifyCommand {
	if r.Create != nil {
		return global.NotifyCreate
	} else if r.Receive != nil {
		return global.NotifyReceive
	} else if r.Read != nil {
		return global.NotifyRead
	} else if r.Delete != nil {
		return global.NotifyDelete
	} else {
		return ""
	}
}

type MentionCreate struct {
	TopicId    int64       `json:"topic_id,omitempty" validate:"required,min=1"`
	UserIds    []int64     `json:"user_ids,omitempty" validate:"required,min=1"`
	SequenceId int64       `json:"sequence_id,omitempty" validate:"required,min=1"`
	Custom     interface{} `json:"custom,omitempty"`
}

type RequestReply struct {
	Create  *ReplyCreate `json:"create,omitempty"`
	Receive *Crud        `json:"receive,omitempty"`
	Read    *Crud        `json:"read,omitempty"`
	Delete  *Crud        `json:"delete,omitempty"`
}

func (r RequestReply) GetNotifyType() global.NotifyCommand {
	if r.Create != nil {
		return global.NotifyCreate
	} else if r.Receive != nil {
		return global.NotifyReceive
	} else if r.Read != nil {
		return global.NotifyRead
	} else if r.Delete != nil {
		return global.NotifyDelete
	} else {
		return ""
	}
}

type ReplyCreate struct {
	TopicId    int64       `json:"topic_id,omitempty" validate:"required,min=1"`
	UserId     int64       `json:"user_id,omitempty" validate:"required,min=1"`
	SequenceId int64       `json:"sequence_id,omitempty" validate:"required,min=1"`
	Custom     interface{} `json:"custom,omitempty"`
}

type Crud struct {
	NotifyId int64 `json:"notify_id,omitempty" validate:"required,min=1"`
}
