package model

import "errors"

var (
	ErrTitleInvalid    = errors.New("invalid title")
	ErrJobTitleInvalid = errors.New("invalid job title")
	ErrJobCtxInvalid   = errors.New("invalid job ctx")
	ErrJobUIDInvalid   = errors.New("invalid job UID")
	ErrJobUIDMismatch  = errors.New("job UID does not match")
	ErrJobOrderInvalid = errors.New("invalid order ID")
	ErrJobNotEditable  = errors.New("job is not waiting for input")
)
