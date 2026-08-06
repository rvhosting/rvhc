package vmgr

import (
	"errors"
)

func Quit(id string) error {
	qmp := Get(id)
	if qmp == nil {
		return errors.New("unknown vm id")
	}

	_, err := qmp.exec(`{"execute":"quit"}`)
	if err != nil {
		return err
	}

	return nil
}
