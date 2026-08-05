package vmgr

import (
	"sync"
)

var vms = map[string]*QMP{}
var mu sync.RWMutex

func Init() {}
