package cli

import (
	"rvhc/daemon"

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
