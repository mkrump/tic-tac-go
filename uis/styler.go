package uis

type Styler interface {
	Style(square int, squareIndex int, marker string) string
}
