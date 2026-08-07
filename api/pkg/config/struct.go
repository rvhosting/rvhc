package config

import (
	"rvhc/daemon"
	"rvhc/protocol"
)

type DaemonConfig struct {
	Listen daemon.ListenConfig `json:"listen"`
	protocol.Auth
}
