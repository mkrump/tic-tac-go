package consoles

import "errors"

var (
	ErrInvalidInput  = errors.New("console: invalid input")
	ErrInvalidOption = errors.New("console: invalid option")
)
