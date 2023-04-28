package utils

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("Record not found")
	UnknownError      = errors.New("UnknownError")
)
