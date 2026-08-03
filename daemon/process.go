package daemon

import (
	"net"
	"rvhc/daemon/internal/fns"
	"rvhc/protocol"
)

func process(conn net.Conn) {
	for {
		var payload protocol.Payload
		if err := protocol.Read(conn, &payload); err != nil {
			protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
			return
		}

		switch payload.Type {

		case "get-limit":
			var names []string
			if err := payload.UnmarshalTo(&names); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			result, err := fns.GetLimit(names)
			if err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
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
