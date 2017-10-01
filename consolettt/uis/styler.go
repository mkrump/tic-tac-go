package uis

import "fmt"

type ColorStyler struct {
}

func (ColorStyler) Style(square int, squareIndex int, marker string) string {
	var red = "\033[1;31m%2s\033[0m"
	var green = "\033[1;32m%2s\033[0m"

	switch square {
	case -1:
		marker = fmt.Sprintf(green, marker)
	case 1:
		marker = fmt.Sprintf(red, marker)
	default:
		marker = fmt.Sprintf("%2d", squareIndex+1)
	}
	return marker
}

type SimpleStyle struct {
}

func (SimpleStyle) Style(square int, squareIndex int, marker string) string {
	switch square {
	case -1, 1:
		marker = fmt.Sprintf("%2s", marker)
	default:
		marker = fmt.Sprintf("%2d", squareIndex)
	}
	return marker
}
