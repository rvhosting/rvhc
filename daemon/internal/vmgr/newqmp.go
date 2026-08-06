package vmgr

import (
	"bufio"
	"os/exec"
)

func NewQMP(cmd *exec.Cmd) (*QMP, error) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	return &QMP{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,

		writer: bufio.NewWriter(stdin),
		reader: bufio.NewReader(stdout),
	}, nil
}
