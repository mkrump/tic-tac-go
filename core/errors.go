package core

import (
	"errors"
)

var (
	ErrSquareOccupied = errors.New("core: square occupied")
	ErrOutOfBounds    = errors.New("core: out bounds")
)
