package vmgr

import (
	"log/slog"
	"os/exec"
)

func Start(id string, qmp *QMP) {
	mu.Lock()
	vms[id] = qmp
	mu.Unlock()

	go func(id string, cmd *exec.Cmd) {
		if err := cmd.Run(); err != nil {
			slog.Error(err.Error(), "id", id)
		}

		mu.Lock()
		delete(vms, id)
		mu.Unlock()
	}(id, qmp.cmd)
}
