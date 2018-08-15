package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var QuoteId string

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a quote from Quotes",
	Long:  `Delete a quote from Quotes given its id.`,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "http://localhost:8080/quotes/" + QuoteId
		client := &http.Client{}
		req, err := http.NewRequest("DELETE", uri, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(os.Stdout, "Quote with ID: '%v' was deleted.\n", QuoteId)
	},
}

func init() {
	clientCmd.AddCommand(delCmd)
	delCmd.Flags().StringVarP(&QuoteId, "quote-id", "q", "", "Id of Quote to remove")
	delCmd.MarkFlagRequired("quote-id")
}
