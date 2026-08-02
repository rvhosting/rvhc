package cli

import (
	"rvhc/daemon"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "rvhc",
	Short: "rvhosting cli",
}

var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start daemon",
	Run:   daemon.Start,
}
