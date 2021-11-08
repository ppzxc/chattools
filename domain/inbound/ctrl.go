package inbound

import (
	"github.com/ppzxc/chattools/common"
)

type Control struct {
	UUID   string         `json:"uuid,omitempty" validate:"required,uuid4"`
	Create *RequestCreate `json:"create,omitempty"`
	Leave  *RequestLeave  `json:"leave,omitempty"`
	Invite *RequestInvite `json:"invite,omitempty"`
	Join   *RequestJoin   `json:"join,omitempty"`
}

func (c Control) GetCtrlType() common.Control {
	if c.Create != nil {
		return common.CtrlCreate
	} else if c.Leave != nil {
		return common.CtrlLeave
	} else if c.Join != nil {
		return common.CtrlJoin
	} else if c.Invite != nil {
		return common.CtrlInvite
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
