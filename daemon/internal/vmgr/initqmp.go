package vmgr

func InitQMP(qmp *QMP) error {
	_, err := qmp.exec(`{"execute":"qmp_capabilities"}`)
	if err != nil {
		return err
	}

	return nil
}
