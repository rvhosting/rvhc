package fns

import (
	"errors"
	"runtime"

	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetLimit(names []string) (map[string]any, error) {
	result := map[string]any{}
	for _, name := range names {
		switch name {

		case "cpu":
			result["cpu"] = runtime.NumCPU()

		case "memory":
			memory, err := mem.VirtualMemory()
			if err != nil {
				return nil, err
			}

			result["memory"] = memory

		case "storage":
			storage, err := disk.Usage("/")
			if err != nil {
				return nil, err
			}

			result["storage"] = storage

		default:
			return nil, errors.New("unknown name")

		}
	}

	return result, nil
}
