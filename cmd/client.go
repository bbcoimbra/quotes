package cmd

import (
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client for Quotes",
	Long:  `client to interact with Quotes API without use curl.`,
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
