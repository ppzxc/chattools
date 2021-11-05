package common

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var Validate *validator.Validate

var (
	WriteTimeOut = 5 * time.Second
)
