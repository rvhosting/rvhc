package cli

import "rvhc/daemon"

func initarg() {
	DaemonCmd.Flags().StringVarP(&daemon.ConfigFile, "config", "c", "daemon.json", "custom daemon config")
}
