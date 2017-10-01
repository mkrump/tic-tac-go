package ui

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var red = "\033[1;31m%2s\033[0m"
var green = "\033[1;32m%2s\033[0m"

func TestSimplerStylerRendersMarkerForOccupiedSquare(t *testing.T) {
	styler := SimpleStyle{}

	styleSquare := styler.Style(-1, 0, "X")

	assert.Equal(t, " X", styleSquare)
}

func TestSimplerStylerRendersSquareForOccupiedSquare(t *testing.T) {
	styler := SimpleStyle{}

	styleSquare := styler.Style(0, 0, "X")

	assert.Equal(t, " 0", styleSquare)
}

func TestColorStylerRendersGreenMarkerForOccupiedSquare(t *testing.T) {
	styler := ColorStyler{}

	styleSquare := styler.Style(-1, 0, "X")
	greenX := fmt.Sprintf(green, "X")

	assert.Equal(t, greenX, styleSquare)
}
func TestColorStylerRendersRedMarkerForOccupiedSquare(t *testing.T) {
	styler := ColorStyler{}

	styleSquare := styler.Style(1, 0, "O")
	redO := fmt.Sprintf(red, "O")

	assert.Equal(t, redO, styleSquare)
}

func TestColorStylerRendersIndexPlusOneForNonOccupiedSquare(t *testing.T) {
	styler := ColorStyler{}

	styleSquare := styler.Style(0, 0, "O")

	assert.Equal(t, " 1", styleSquare)
}
