package daemon

import (
	"net"
	"rvhc/daemon/internal/dataimg"
	"rvhc/daemon/internal/db"
	"rvhc/daemon/internal/fns"
	"rvhc/daemon/internal/vmgr"
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

		case "get-vm-info":
			var vmID string
			if err := payload.UnmarshalTo(&vmID); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			var vm db.VM
			result := db.GetInfo(vmID, &vm)
			if result.Error != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: result.Error.Error()})
				return
			}

			if err := protocol.Write(conn, vm); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			protocol.Write(conn, protocol.Status{Success: true})

		case "start":
			var vmID string
			if err := payload.UnmarshalTo(&vmID); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			var vm db.VM
			result := db.GetInfo(vmID, &vm)
			if result.Error != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: result.Error.Error()})
				return
			}

			cmd := vmgr.New(vm)
			qmp, err := vmgr.NewQMP(cmd)
			if err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			vmgr.Start(vm.ID, qmp)
			if err := vmgr.InitQMP(qmp); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			markResult := db.Mark(vm.ID, true)
			if markResult.Error != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: markResult.Error.Error()})
				return
			}

			protocol.Write(conn, protocol.Status{Success: true})

		case "stop":
			var vmID string
			if err := payload.UnmarshalTo(&vmID); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			if err := vmgr.Quit(vmID); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			markResult := db.Mark(vmID, false)
			if markResult.Error != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: markResult.Error.Error()})
				return
			}

			protocol.Write(conn, protocol.Status{Success: true})

		case "get-ports":
			var vmID string
			if err := payload.UnmarshalTo(&vmID); err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			portMap, err := vmgr.GetPorts(vmID)
			if err != nil {
				protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
				return
			}

			if err := protocol.Write(conn, portMap); err != nil {
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
