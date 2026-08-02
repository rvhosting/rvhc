package cli

func initCmd() {
	rootCmd.AddCommand(daemonCmd)
	rootCmd.AddCommand(downloadCmd)
}
