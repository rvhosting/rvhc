package daemon

import (
	"net"
	"rvhc/protocol"
)

func auth(conn net.Conn) bool {
	var userAuth protocol.Auth
	if err := protocol.Read(conn, &userAuth); err != nil {
		protocol.Write(conn, protocol.Status{Success: false, Message: err.Error()})
		return false
	}

	success := userAuth == config.Auth
	if !success {
		protocol.Write(conn, protocol.Status{Success: false, Message: "auth failed"})
		return false
	}

	protocol.Write(conn, protocol.Status{Success: true})
	return true
}
