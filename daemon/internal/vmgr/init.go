package vmgr

import (
	"os/exec"
	"sync"
)

var vms = map[string]*exec.Cmd{}
var mu sync.RWMutex

func Init() {}
