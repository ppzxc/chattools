package inbound

import (
	"github.com/ppzxc/chattools/common"
)

type Authentication struct {
	UUID     string           `json:"uuid,omitempty" validate:"required,uuid4"`
	Login    *RequestLogin    `json:"login,omitempty"`
	Register *RequestRegister `json:"register,omitempty"`
	Logout   *RequestLogout   `json:"logout,omitempty"`
	Token    *RequestToken    `json:"token,omitempty"`
}

func (a Authentication) GetAuthType() common.Authentication {
	if a.Login != nil {
		return common.AuthLogin
	} else if a.Logout != nil {
		return common.AuthLogout
	} else if a.Token != nil {
		return common.AuthToken
	} else if a.Register != nil {
		return common.AuthRegister
	}
	return ""
}

type RequestLogin struct {
	Using      string     `json:"using" validate:"required,oneof=id anonymous token rotary"`
	User       *LoginUser `json:"user,omitempty" validate:"required_if=Using id"`            // for id
	Token      string     `json:"token,omitempty" validate:"required_if=Using token"`        // for token
	UserId     int64      `json:"user_id,omitempty" validate:"required_if=Using rotary"`     // for rotary
	UserName   string     `json:"username,omitempty" validate:"required_if=Using anonymous"` // for anonymous
	DeviceInfo DeviceInfo `json:"device_info" validate:"required"`
}

type RequestRegister struct {
	Name       string     `json:"name" validate:"required,min=1,max=40"`
	Email      string     `json:"email" validate:"required,email"`
	Password   string     `json:"password" validate:"required,min=10,max=40"`
	DeviceInfo DeviceInfo `json:"device_info" validate:"required"`
}

type RequestLogout struct {
}

type RequestToken struct {
}

//type RegisterUser struct {
//	Name     string `json:"name" validate:"required,min=1,max=40"`
//	Email    string `json:"email" validate:"required,email"`
//	Password string `json:"password" validate:"required,min=10,max=40"`
//}
//

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10,max=40"`
}
