package tttuis

import (
	"errors"
)

var (
	ErrInvalidInput  = errors.New("ui: invalid input")
	ErrInvalidOption = errors.New("ui: invalid option")
)
