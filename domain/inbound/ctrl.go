package inbound

import (
	"github.com/ppzxc/chattools/common/global"
)

type Control struct {
	UUID   string         `json:"uuid,omitempty" validate:"required,uuid4"`
	Create *RequestCreate `json:"create,omitempty"`
	Leave  *RequestLeave  `json:"leave,omitempty"`
	Invite *RequestInvite `json:"invite,omitempty"`
	Join   *RequestJoin   `json:"join,omitempty"`
}

func (c Control) GetCtrlType() global.Control {
	if c.Create != nil {
		return global.CtrlCreate
	} else if c.Leave != nil {
		return global.CtrlLeave
	} else if c.Join != nil {
		return global.CtrlJoin
	} else if c.Invite != nil {
		return global.CtrlInvite
	} else {
		return ""
	}
}

type RequestCreate struct {
	Name string `json:"name,omitempty" validate:"required"`
}
type RequestLeave struct {
	TopicId int64 `json:"topic_id,omitempty" validate:"required"`
}
type RequestInvite struct {
	TopicId int64 `json:"topic_id,omitempty" validate:"required"`
	UserId  int64 `json:"user_id,omitempty" validate:"required"`
}
type RequestJoin struct {
	TopicId int64 `json:"topic_id,omitempty" validate:"required"`
}
