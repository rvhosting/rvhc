package vmgr

import (
	"log"
	"rvhc/daemon/internal/db"
	"sync"
)

var vms = map[string]*QMP{}
var mu sync.RWMutex

func Init() {
	var runningVMs []db.VM
	result := db.GetRunningVMs(&runningVMs)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	for _, vm := range runningVMs {
		cmd := New(vm)
		qmp, err := NewQMP(cmd)
		if err != nil {
			log.Fatalln(err)
		}

		if err := Start(vm.ID, qmp); err != nil {
			log.Fatalln(err)
		}

		if err := InitQMP(qmp); err != nil {
			log.Fatalln(err)
		}
	}
}
