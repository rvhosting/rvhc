package vmgr

func Get(id string) *QMP {
	mu.RLock()
	defer mu.RUnlock()

	qmp, ok := vms[id]
	if !ok {
		return nil
	}

	return qmp
}
