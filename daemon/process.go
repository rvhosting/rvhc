package daemon

import (
	"net"
	"runtime"
	"rvhc/protocol"

	"github.com/shirou/gopsutil/v4/mem"
)

func process(conn net.Conn) {
	for {
		var payload protocol.Payload
		if err := protocol.Read(conn, &payload); err != nil {
			protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
			return
		}

		switch payload.Type {

		case "limit":
			var names []string
			if err := payload.UnmarshalTo(&names); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			result := map[string]any{}
			for _, name := range names {

				if name == "cpu" {
					result["cpu"] = runtime.NumCPU()
					continue
				}

				if name == "memory" {
					memory, err := mem.VirtualMemory()
					if err != nil {
						protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
						return
					}

					result["memory"] = memory
					continue
				}
			}

			if err := protocol.Write(conn, result); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			protocol.Write(conn, protocol.Status{Success: true})

		default:
			protocol.Write(conn, protocol.Status{Success: false, Message: "unknown type"})
			return

		}
	}
}
