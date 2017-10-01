package uis

import (
	"errors"
)

var (
	ErrInvalidInput  = errors.New("uis: invalid input")
	ErrInvalidOption = errors.New("uis: invalid option")
)
