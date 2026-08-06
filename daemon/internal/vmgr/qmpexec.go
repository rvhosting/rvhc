package vmgr

import (
	"log"
	"strings"
)

func (q *QMP) exec(cmd string) (string, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	log.Println("exec qmp command", cmd)
	if _, err := q.writer.WriteString(cmd + "\n"); err != nil {
		return "", err
	}

	log.Println("wait command", cmd, "send")
	if err := q.writer.Flush(); err != nil {
		return "", err
	}

	log.Println("wait command", cmd, "response")
	for {
		output, err := q.reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		if strings.Contains(output, `"return"`) || strings.Contains(output, `"error"`) {
			log.Println("command", cmd, "executed")
			return output, nil
		}
	}
}
