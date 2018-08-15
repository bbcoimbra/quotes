package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type QuoteMsg struct {
	Author, Text string
}

var Author, Text string

var addCmd = &cobra.Command{
	Use:   "add --text 'your quote text here' --author 'Author Name'",
	Short: "Add a quote to quotes database",
	Long: `Add a quhote to quotes database. Flags --text and --author are required.
An example can be as below:

	quotes add --text "A vingança nunca é plena. Mata a alma e a envenena."`,
	Run: func(cmd *cobra.Command, args []string) {
		quote := &QuoteMsg{Text: Text, Author: Author}
		json, err := json.Marshal(quote)
		if err != nil {
			log.Fatal(err)
		}
		buf := bytes.NewBuffer(json)
		resp, err := http.Post("http://localhost:8080/quotes/", "application/json", buf)
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
	clientCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Text, "text", "t", "", "Text of the quote")
	addCmd.Flags().StringVarP(&Author, "author", "a", "", "Author of the quote")
	addCmd.MarkFlagRequired("text")
	addCmd.MarkFlagRequired("author")
}
