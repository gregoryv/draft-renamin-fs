package oserror

import "errors"

var (
	ErrInvalid    = errors.New("ErrInvalid")
	ErrPermission = errors.New("ErrPermission")
	ErrExist      = errors.New("ErrExist")
	ErrNotExist   = errors.New("ErrNotExist")
	ErrClosed     = errors.New("ErrClosed")
)
