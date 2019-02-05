package vlc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerConfig_ActualExePath(t *testing.T) {
	cfg := PlayerConfig{}
	assert.Equal(t, cfg.defaultExePath(), cfg.ActualExePath())

	cfg = PlayerConfig{ExePath: "/1/2/3"}
	assert.Equal(t, "/1/2/3", cfg.ActualExePath())
}

func TestPlayerConfig_ActualTCPPort(t *testing.T) {
	cfg := PlayerConfig{}
	assert.Equal(t, defaultTCPPort, cfg.ActualTCPPort())

	cfg = PlayerConfig{TCPPort: 12345}
	assert.Equal(t, 12345, cfg.ActualTCPPort())
}
