package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ServerURL string
	QuotePath string
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client for Quotes",
	Long:  `client to interact with Quotes API without use curl.`,
}

func init() {
	rootCmd.AddCommand(clientCmd)
	ServerURL = viper.GetString("client.api_endpoint")
	QuotePath = "/quotes/"
}
