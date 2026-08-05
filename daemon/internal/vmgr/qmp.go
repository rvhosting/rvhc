package vmgr

import (
	"io"
	"os/exec"
)

type QMP struct {
	cmd    *exec.Cmd
	stdin  io.Writer
	stdout io.Reader
}

func (q *QMP) Read(p []byte) (n int, err error) {
	if q.stdout == nil {
		return 0, io.EOF
	}

	return q.stdout.Read(p)
}

func (q *QMP) Write(p []byte) (n int, err error) {
	if q.stdin == nil {
		return 0, io.ErrClosedPipe
	}

	return q.stdin.Write(p)
}

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
	}, nil
}
