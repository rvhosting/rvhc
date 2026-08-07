package api

import (
	"rvhc/api/config"
)

var ConfigFile string
var DaemonName string

func Init() {
	config.Load(ConfigFile, DaemonName)
}
