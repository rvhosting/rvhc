package cli

import (
	"rvhc/daemon"
	"rvhc/downloader"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rvhc",
	Short: "rvhosting cli",
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start daemon",
	Run: func(cmd *cobra.Command, args []string) {
		daemon.Start()
	},
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download minihttpsys",
	Run: func(cmd *cobra.Command, args []string) {
		downloader.Download()
	},
}
