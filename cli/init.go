package cli

import "os"

func init() {
	initarg()

	RootCmd.AddCommand(DaemonCmd)

	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
