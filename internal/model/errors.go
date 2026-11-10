package model

import "errors"

var (
	ErrTitleInvalid    = errors.New("invalid title")
	ErrJobTitleInvalid = errors.New("invalid job title")
	ErrJobCtxInvalid   = errors.New("invalid job ctx")
)
