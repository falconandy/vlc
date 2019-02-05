package vlc

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	startupDelay = time.Second * 2
)

type Duration int

type Player struct {
	exePath string
	tcpPort int
	trackRe *regexp.Regexp

	debugMode bool
	conn      *tcpConnection
	commands  chan<- *command
}

func NewPlayer(cfg *PlayerConfig) *Player {
	if cfg == nil {
		cfg = &PlayerConfig{}
	}
	exePath, tcpPort := cfg.ActualExePath(), cfg.ActualTCPPort()

	return &Player{
		exePath: exePath,
		tcpPort: tcpPort,
		trackRe: regexp.MustCompile(`^\| (\d+) - (.*)$`),
	}
}

func (p *Player) SetDebugMode(debugMode bool) *Player {
	p.debugMode = debugMode
	return p
}

func (p *Player) Start() error {
	args := []string{
		"--extraintf=rc",
		fmt.Sprintf("--rc-host=%s:%d", "localhost", p.tcpPort),
		"--one-instance",
	}

	// TODO: specific to a VLC version?
	if runtime.GOOS == "windows" {
		args = append(args, "--rc-quiet")
	}

	cmd := exec.Command(p.exePath, args...)
	err := cmd.Start()
	if err != nil {
		return err
	}

	time.Sleep(startupDelay)

	p.conn = newTCPConnection(p.tcpPort, p.debugMode)
	err = p.conn.Open()
	if err != nil {
		return err
	}

	commands := make(chan *command)
	p.commands = commands
	go p.conn.run(commands)

	return nil
}

func (p *Player) Shutdown() error {
	_, err := p.execCommand(p.conn.version.shutdownCommand)
	return err
}

func (p *Player) Play(videoPath string) error {
	_, err := p.execCommand(fmt.Sprintf(`add %s`, videoPath))
	return err
}

func (p *Player) IsPlaying() (bool, error) {
	output, err := p.execCommand("is_playing")
	if err != nil {
		return false, err
	}
	return len(output) != 0 && output[0] == "1", nil
}

func (p *Player) Pause() error {
	_, err := p.execCommand("pause")
	return err
}

func (p *Player) Stop() error {
	_, err := p.execCommand("stop")
	return err
}

func (p *Player) Length() (Duration, error) {
	output, err := p.execCommand("get_length")
	if err != nil {
		return -1, err
	}
	if len(output) == 0 {
		return -1, errors.New("unexpected empty output of a 'get_length' command")
	}
	length, err := strconv.Atoi(output[0])
	if err != nil {
		return -1, fmt.Errorf("can't convert '%s' to a number: %v", output, err)
	}
	return Duration(length), nil
}

func (p *Player) Position() (Duration, error) {
	output, err := p.execCommand("get_time")
	if err != nil {
		return -1, err
	}
	if len(output) == 0 {
		return -1, nil
	}
	position, err := strconv.Atoi(output[0])
	if err != nil {
		return -1, fmt.Errorf("can't convert '%s' to a number: %v", output, err)
	}
	return Duration(position), nil
}

func (p *Player) Seek(position Duration) error {
	_, err := p.execCommand(fmt.Sprintf("seek %d", position))
	return err
}

func (p *Player) SpeedSlower() error {
	_, err := p.execCommand("slower")
	return err
}

func (p *Player) SpeedFaster() error {
	_, err := p.execCommand("faster")
	return err
}

func (p *Player) SpeedNormal() error {
	_, err := p.execCommand("normal")
	return err
}

func (p *Player) AudioTracks() (tracks map[int]string, activeTrack int, err error) {
	output, err := p.execCommand("atrack")
	if err != nil {
		return nil, -1, err
	}
	if len(output) == 0 {
		return nil, -1, errors.New("unexpected empty output of a 'atrack' command")
	}

	tracks, activeTrack = p.parseTracks(output)
	return tracks, activeTrack, nil
}

func (p *Player) SetAudioTrack(track int) error {
	_, err := p.execCommand(fmt.Sprintf("atrack %d", track))
	return err
}

func (p *Player) SubtitleTrack() (tracks map[int]string, activeTrack int, err error) {
	output, err := p.execCommand("strack")
	if err != nil {
		return nil, -1, err
	}
	if len(output) == 0 {
		return nil, -1, errors.New("unexpected empty output of a 'strack' command")
	}

	tracks, activeTrack = p.parseTracks(output)
	return tracks, activeTrack, nil
}

func (p *Player) SetSubtitleTrack(track int) error {
	_, err := p.execCommand(fmt.Sprintf("strack %d", track))
	return err
}

func (p *Player) execCommand(cmd string) (output []string, err error) {
	c := newCommand(cmd)
	p.commands <- c
	result := <-c.result
	return result.output, result.err
}

func (p *Player) parseTracks(output []string) (tracks map[int]string, activeTrack int) {
	tracks = make(map[int]string, len(output))
	activeTrack = -1
	for _, line := range output {
		match := p.trackRe.FindStringSubmatch(line)
		if match == nil {
			continue
		}

		trackIndex, _ := strconv.Atoi(match[1])
		trackTitle := match[2]
		if strings.HasSuffix(trackTitle, "*") {
			activeTrack = trackIndex
			trackTitle = strings.TrimSpace(strings.TrimSuffix(trackTitle, "*"))
		}
		tracks[trackIndex] = trackTitle
	}
	return tracks, activeTrack
}
