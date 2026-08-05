package vmgr

import (
	"bufio"
	"errors"
	"io"
)

func Quit(id string) error {
	qmp := Get(id)
	if qmp == nil {
		return errors.New("unknown vm id")
	}

	writer := bufio.NewWriter(qmp)
	if _, err := writer.WriteString(`{"execute":"quit"}` + "\n"); err != nil {
		return err
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	if _, err := io.ReadAll(qmp); err != nil {
		return err
	}

	return nil
}
