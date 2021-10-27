package common

import "errors"

var (
	ErrValidateNotContainsRoutingObjectInAuth = errors.New("request object is not contains in auth, [login, logout, register]")
	ErrValidateNotContainsRequestObject       = errors.New("request object is not contains, [auth, ctrl, meta, msg, notify]")
)
