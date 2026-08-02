package daemon

import (
	"rvhc/protocol"
)

type ListenConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type daemonConfig struct {
	Listen ListenConfig `json:"listen"`
	protocol.Auth
}
