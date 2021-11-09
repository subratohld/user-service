package main

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the service",
	Long:  "Starts the service",
	Run: func(cmd *cobra.Command, args []string) {
		startApp()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
