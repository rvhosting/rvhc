package vmgr

import (
	"bufio"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func GetPorts(id string) (map[int]int, error) {
	qmp := Get(id)
	if qmp == nil {
		return nil, errors.New("unknown vm id")
	}

	log.Println("execute human-monitor-command")
	writer := bufio.NewWriter(qmp)
	if _, err := writer.WriteString(`{"execute":"human-monitor-command","arguments":{"command-line":"info usernet"}}` + "\n"); err != nil {
		return nil, err
	}

	log.Println("wait human-monitor-command send")
	if err := writer.Flush(); err != nil {
		return nil, err
	}

	log.Println("wait human-monitor-command response")
	reader := bufio.NewReader(qmp)
	var rawOutput string

	for {
		raw, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		if strings.Contains(raw, `"return"`) {
			rawOutput = raw
			break
		}
	}

	re := regexp.MustCompile(`127\.0\.0\.1\s+(\d+)\s+[\d\.]+\s+(\d+)`)
	matches := re.FindAllStringSubmatch(rawOutput, -1)
	portMap := map[int]int{}

	for _, match := range matches {
		if len(match) == 3 {
			hostPort, _ := strconv.Atoi(match[1])
			guestPort, _ := strconv.Atoi(match[2])

			if guestPort == 22 || guestPort == 80 {
				portMap[guestPort] = hostPort
			}
		}
	}

	return portMap, nil
}
