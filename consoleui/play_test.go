package consoleui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlay(t *testing.T) {
	expected := "\n " +
		"X | O | X \n" +
		"   | O |  \n" +
		"   | O |  \n"
	assert.Equal(t, expected, Play(), "")
}
