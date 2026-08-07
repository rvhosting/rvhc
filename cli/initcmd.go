package cli

func initCmd() {
	rootCmd.AddCommand(daemonCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(apiCmd)
}
