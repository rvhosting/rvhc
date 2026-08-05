package vmgr

import (
	"log"
	"log/slog"
	"os/exec"
)

func Start(id string, qmp *QMP) {
	mu.Lock()
	vms[id] = qmp
	mu.Unlock()

	go func(id string, cmd *exec.Cmd) {
		log.Println("exec", cmd.String())

		if err := cmd.Run(); err != nil {
			slog.Error(err.Error(), "id", id)
		}

		mu.Lock()
		delete(vms, id)
		mu.Unlock()

		log.Println("vm", id, "stop")
	}(id, qmp.cmd)
}
