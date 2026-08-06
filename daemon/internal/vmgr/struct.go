package vmgr

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

type QMP struct {
	cmd    *exec.Cmd
	stdin  io.Writer
	stdout io.Reader

	writer *bufio.Writer
	reader *bufio.Reader
	mu     sync.Mutex
}
