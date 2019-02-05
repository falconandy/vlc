package vlc

type commandResult struct {
	output []string
	err    error
}

type command struct {
	cmd    string
	result chan *commandResult
}

func newCommand(cmd string) *command {
	return &command{
		cmd:    cmd,
		result: make(chan *commandResult, 1),
	}
}

func (c *command) String() string {
	return c.cmd
}
