package inbound

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ppzxc/chattools/common"
	"github.com/ppzxc/chattools/types"
)

type Root struct {
	Auth   *Authentication `json:"auth,omitempty"`
	Ctrl   *Control        `json:"ctrl,omitempty"`
	Msg    *Message        `json:"msg,omitempty"`
	Meta   *Meta           `json:"meta,omitempty"`
	Notify *Notify         `json:"notify,omitempty"`
	File   *File           `json:"file,omitempty"`
	Ping   *Ping           `json:"ping,omitempty"`
}

func (r Root) GetInboundType() common.Bound {
	if r.Auth != nil {
		return common.BoundAuthentication
	} else if r.Ctrl != nil {
		return common.BoundControl
	} else if r.Msg != nil {
		return common.BoundMessage
	} else if r.Meta != nil {
		return common.BoundMeta
	} else if r.Notify != nil {
		return common.BoundNotify
	} else if r.File != nil {
		return common.BoundFile
	} else if r.Ping != nil {
		return common.BoundPing
	} else {
		return common.BoundEtc
	}
}

func (r Root) Validate() error {
	if r.Auth != nil {
		if err := common.Validate.Struct(r.Auth); err != nil {
			return Change(err)
		}
		return nil
		//} else {
		//	if r.Auth.Login != nil {
		//		return Change(common.Validate.Struct(r.Auth.Login))
		//	} else if r.Auth.Logout != nil {
		//		return Change(common.Validate.Struct(r.Auth.Logout))
		//	} else if r.Auth.Token != nil {
		//		return Change(common.Validate.Struct(r.Auth.Token))
		//	} else if r.Auth.Register != nil {
		//		return Change(common.Validate.Struct(r.Auth.Register))
		//	} else {
		//		return types.ErrValidateNotContainsRoutingObjectInAuth
		//	}
		//}
	} else if r.Ctrl != nil {
		if err := common.Validate.Struct(r.Ctrl); err != nil {
			return Change(err)
		}
		return nil
	} else if r.Msg != nil {
		if err := common.Validate.Struct(r.Msg); err != nil {
			return Change(err)
		}
		return nil
	} else if r.Meta != nil {
		if err := common.Validate.Struct(r.Meta); err != nil {
			return Change(err)
		}
		return nil
	} else if r.Notify != nil {
		if err := common.Validate.Struct(r.Notify); err != nil {
			return Change(err)
		}
		return nil
	} else if r.File != nil {
		if err := common.Validate.Struct(r.File); err != nil {
			return Change(err)
		}
		return nil
	} else if r.Ping != nil {
		if err := common.Validate.Struct(r.Ping); err != nil {
			return Change(err)
		}
		return nil
	} else {
		return types.ErrValidateNotContainsRequestObject
	}
}

func Change(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	var s string
	for _, err := range err.(validator.ValidationErrors) {
		if len(err.Param()) <= 0 && err.Value() != nil {
			s = s + err.Error() + "\r\n"
		} else {
			s = s + err.Error() + fmt.Sprintf(" [Expect:%v, Actual:%v]", err.Param(), err.Value()) + "\r\n"
		}
	}
	return errors.New(s)
}

type File struct {
	UUID string `json:"uuid,omitempty" validate:"required,uuid4"`
}
type Ping struct {
}
