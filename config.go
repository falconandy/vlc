package vlc

import (
	"runtime"
)

const (
	defaultTCPPort        = 2019
	defaultVLCPath        = "vlc"
	defaultWindowsVLCPath = `C:\Program Files (x86)\VideoLAN\VLC\vlc.exe`
)

type PlayerConfig struct {
	ExePath string
	TCPPort int
}

func (c PlayerConfig) ActualExePath() string {
	if c.ExePath != "" {
		return c.ExePath
	}

	return c.defaultExePath()
}

func (c PlayerConfig) ActualTCPPort() int {
	if c.TCPPort != 0 {
		return c.TCPPort
	}

	return defaultTCPPort
}

func (c PlayerConfig) defaultExePath() string {
	switch runtime.GOOS {
	case "windows":
		return defaultWindowsVLCPath
	default:
		return defaultVLCPath
	}
}
