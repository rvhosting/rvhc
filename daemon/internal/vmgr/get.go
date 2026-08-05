package vmgr

import (
	"os/exec"
)

func Get(id string) *exec.Cmd {
	mu.RLock()
	defer mu.RUnlock()

	cmd, ok := vms[id]
	if !ok {
		return nil
	}

	return cmd
}
