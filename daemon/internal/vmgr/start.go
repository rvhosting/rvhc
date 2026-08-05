package vmgr

import (
	"log/slog"
	"os/exec"
)

func Start(id string, cmd *exec.Cmd) {
	mu.Lock()
	vms[id] = cmd
	mu.Unlock()

	go func(id string, cmd *exec.Cmd) {
		if err := cmd.Run(); err != nil {
			slog.Error(err.Error(), "id", id)
		}

		mu.Lock()
		delete(vms, id)
		mu.Unlock()
	}(id, cmd)
}
