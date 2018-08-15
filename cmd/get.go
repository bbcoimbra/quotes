package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Quote by its id",
	Long:  `Get Quote by its id.`,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "http://localhost:8080/quotes/" + QuoteId
		resp, err := http.Get(uri)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()

		quoteResponse, err := ioutil.ReadAll(resp.Body)

		fmt.Fprintf(os.Stdout, "%v\n", string(quoteResponse))
	},
}

func init() {
	clientCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&QuoteId, "quote-id", "q", "", "Id of the Quote to get")
	getCmd.MarkFlagRequired("quote-id")
}
