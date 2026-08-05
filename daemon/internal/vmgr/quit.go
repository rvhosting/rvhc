package vmgr

import (
	"bufio"
	"errors"
	"io"
	"log"
)

func Quit(id string) error {
	qmp := Get(id)
	if qmp == nil {
		return errors.New("unknown vm id")
	}

	log.Println("execute quit command")
	writer := bufio.NewWriter(qmp)
	if _, err := writer.WriteString(`{"execute":"quit"}` + "\n"); err != nil {
		return err
	}

	log.Println("wait quit command send")
	if err := writer.Flush(); err != nil {
		return err
	}

	log.Println("wait quit command response")
	if _, err := io.ReadAll(qmp); err != nil {
		return err
	}

	return nil
}
