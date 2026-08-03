package daemon

import (
	"log/slog"
	"net"
	"rvhc/protocol"
)

func auth(conn net.Conn) bool {
	var userAuth protocol.Auth
	if err := protocol.Read(conn, &userAuth); err != nil {
		slog.Error(err.Error())
		return false
	}

	success := userAuth == config.Auth
	protocol.Write(conn, protocol.Status{Success: success})
	return success
}
