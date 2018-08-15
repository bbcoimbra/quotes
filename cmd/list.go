package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Quotes",
	Long:  `List all Quotes. JSON formatted output`,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "http://localhost:8080/quotes/list"
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
	clientCmd.AddCommand(listCmd)
}
