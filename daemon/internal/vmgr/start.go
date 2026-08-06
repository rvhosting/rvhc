package vmgr

import (
	"log"
	"log/slog"
	"os/exec"
	"rvhc/daemon/internal/db"
)

func Start(id string, qmp *QMP) error {
	result := db.Mark(id, true)
	if result.Error != nil {
		return result.Error
	}

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

		result := db.Mark(id, false)
		if result.Error != nil {
			slog.Error(result.Error.Error())
		}
	}(id, qmp.cmd)

	return nil
}
