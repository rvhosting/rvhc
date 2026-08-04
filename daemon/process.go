package daemon

import (
	"net"
	"rvhc/daemon/internal/dataimg"
	"rvhc/daemon/internal/db"
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

		case "create":
			var vm db.VM
			if err := payload.UnmarshalTo(&vm); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			vm.Running = false
			result := db.Create(&vm)
			if result.Error != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: result.Error.Error()})
				return
			}

			if err := dataimg.Create(vm.ID, vm.DataSize); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			if err := protocol.Write(conn, vm); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			protocol.Write(conn, protocol.Status{Success: true})

		case "delete":
			var vmID string
			if err := payload.UnmarshalTo(&vmID); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			result := db.Delete(vmID)
			if result.Error != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: result.Error.Error()})
				return
			}

			if err := dataimg.Delete(vmID); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			protocol.Write(conn, protocol.Status{Success: true})

		case "get-vms":
			var vmIDs []string
			result := db.GetIDs(&vmIDs)
			if result.Error != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: result.Error.Error()})
				return
			}

			if err := protocol.Write(conn, vmIDs); err != nil {
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
