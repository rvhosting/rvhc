package vmgr

import (
	"errors"
	"regexp"
	"strconv"
)

func GetPorts(id string) (map[int]int, error) {
	qmp := Get(id)
	if qmp == nil {
		return nil, errors.New("unknown vm id")
	}

	output, err := qmp.exec(`{"execute":"human-monitor-command","arguments":{"command-line":"info usernet"}}`)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`127\.0\.0\.1\s+(\d+)\s+[\d\.]+\s+(\d+)`)
	matches := re.FindAllStringSubmatch(output, -1)
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
