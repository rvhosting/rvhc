package api

import (
	"rvhc/api/pkg/config"
	"rvhc/api/pkg/dconn"
)

var ConfigFile string
var DaemonName string

func Init() {
	config.Load(ConfigFile, DaemonName)
	dconn.Open(config.Daemon)
}
