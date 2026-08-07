package vmgr

func Get(id string) *QMP {
	mu.RLock()
	defer mu.RUnlock()

	qmp, exist := vms[id]
	if !exist {
		return nil
	}

	return qmp
}
