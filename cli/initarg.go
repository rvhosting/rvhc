package cli

import (
	"rvhc/daemon"
	"rvhc/downloader"
)

func initArg() {
	daemonCmd.Flags().StringVarP(&daemon.ConfigFile, "config", "c", "daemon.json", "custom daemon config")
	downloadCmd.Flags().IntVar(&downloader.MaxRetry, "max-retry", 3, "max download retry")
}
