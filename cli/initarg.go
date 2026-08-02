package cli

import "rvhc/daemon"

func initArg() {
	daemonCmd.Flags().StringVarP(&daemon.ConfigFile, "config", "c", "daemon.json", "custom daemon config")
}
