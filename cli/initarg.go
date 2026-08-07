package cli

import (
	"rvhc/api"
	"rvhc/daemon"
	"rvhc/downloader"
)

func initArg() {
	daemonCmd.Flags().StringVarP(&daemon.ConfigFile, "config", "c", "daemon.json", "custom daemon config")
	downloadCmd.Flags().IntVar(&downloader.MaxRetry, "max-retry", 3, "max download retry")

	apiCmd.PersistentFlags().StringVarP(&api.ConfigFile, "config", "c", "api.json", "custom api config")
	apiCmd.PersistentFlags().StringVarP(&api.DaemonName, "daemon", "d", "default", "custom daemon name")
}
