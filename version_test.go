package vlc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_Get(t *testing.T) {
	factory := newVersionFactory()

	assert.Equal(t, "4.0.0", factory.Get("5.0.0").version.String())
	assert.Equal(t, "4.0.0", factory.Get("4.0.0").version.String())
	assert.Equal(t, "0.0.0", factory.Get("3.9.0").version.String())
	assert.Equal(t, "0.0.0", factory.Get("0.0.0").version.String())
	assert.Equal(t, "0.0.0", factory.Get("").version.String())
	assert.Equal(t, "0.0.0", factory.Get("bad").version.String())
}

func TestVersion_Detect(t *testing.T) {
	factory := newVersionFactory()

	text := []string{
		"Command Line Interface initialized. Type 'help' for help.",
		"VLC media player 4.1.3",
		"Command Line Interface initialized. Type 'help' for help.",
	}

	assert.Equal(t, "4.0.0", factory.Detect(text).version.String())
}

func TestVersion_DetectFailed(t *testing.T) {
	factory := newVersionFactory()

	text := []string{
		"Command Line Interface initialized. Type 'help' for help.",
		"VLC 4.1.3",
		"Command Line Interface initialized. Type 'help' for help.",
	}

	assert.Equal(t, "0.0.0", factory.Detect(text).version.String())
}
