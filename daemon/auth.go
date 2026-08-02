package daemon

import (
	"encoding/json"
	"log/slog"
	"net"
	"rvhc/protocol"
)

func auth(conn net.Conn) bool {
	var userAuth protocol.Auth
	if err := json.NewDecoder(conn).Decode(&userAuth); err != nil {
		slog.Error(err.Error())
		return false
	}

	return userAuth == config.Auth
}
