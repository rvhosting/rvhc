package cli

import "os"

func Init() {
	initArg()
	initCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
