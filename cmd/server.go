package cmd

import (
	"log"
	"net/http"

	"github.com/bbcoimbra/quotes/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server starts quotes server",
	Long: `Server starts quotes server in default 8080 TCP port. There is no
configuration for this setting yet`,
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/quotes/", handler.Quotes)
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
