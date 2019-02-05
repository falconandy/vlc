package vlc

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	initialReadTimeout     = time.Millisecond * 500
	nextReadTimeout        = time.Millisecond * 200
	defaultShutdownCommand = "shutdown"
	windowsShutdownCommand = "quit"
)

type tcpConnection struct {
	port            int
	debugMode       bool
	promptRe        *regexp.Regexp
	shutdownCommand string

	conn       net.Conn
	connReader *bufio.Reader
}

func newTCPConnection(port int, debugMode bool) *tcpConnection {
	return &tcpConnection{
		port:      port,
		debugMode: debugMode,
		promptRe:  regexp.MustCompile(`(>\s+)+`),
	}
}

func (c *tcpConnection) Open() error {
	var err error
	c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", c.port))
	if err != nil {
		return err
	}
	c.connReader = bufio.NewReader(c.conn)

	_, err = c.execCommand("help", false)
	if err != nil {
		return err
	}

	c.shutdownCommand = defaultShutdownCommand
	if runtime.GOOS == "windows" {
		c.shutdownCommand = windowsShutdownCommand
	}

	return nil
}

func (c *tcpConnection) run(commands <-chan *command) {
	defer func() {
		_ = c.conn.Close()
		c.connReader = nil
		c.conn = nil
	}()

	for cmd := range commands {
		output, err := c.execCommand(cmd.cmd, c.debugMode)
		cmd.result <- &commandResult{output: output, err: err}
		close(cmd.result)

		if cmd.cmd == c.shutdownCommand {
			break
		}
	}
}

func (c *tcpConnection) execCommand(command string, debug bool) ([]string, error) {
	if debug {
		fmt.Printf("CMD: %s\n", command)
	}

	_, err := fmt.Fprintln(c.conn, command)
	if err != nil {
		return nil, err
	}

	var output []string
	readTimeout := initialReadTimeout
	for {
		err := c.conn.SetReadDeadline(time.Now().Add(readTimeout))
		if err != nil {
			fmt.Printf("can't set read deadline for a VLC connection: %v\n", err)
			break
		}

		line, err := c.connReader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(c.promptRe.ReplaceAllLiteralString(line, ""))

		if strings.HasPrefix(line, "status change:") {
			continue
		}

		if debug {
			fmt.Println("    ", line)
		}

		output = append(output, line)
		readTimeout = nextReadTimeout
	}
	return output, nil
}
