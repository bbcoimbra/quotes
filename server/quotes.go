package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	types "github.com/bbcoimbra/quotes/types"
)

var (
	quotes         []types.Quote
	currentQuoteId int
)

func Quotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		strQuoteId := r.URL.Path[len("/quotes/"):]
		if strQuoteId == "list" {
			quotes = types.Quotes()
			response, _ := json.Marshal(quotes)
			w.Write(response)
			return
		}
		quoteId, err := strconv.Atoi(strQuoteId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		quote := types.GetQuote(quoteId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		response, _ := json.Marshal(quote)
		w.Write(response)
	case http.MethodPost:
		quote := &types.Quote{}
		err := json.NewDecoder(r.Body).Decode(quote)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		types.AddQuote(quote)
		result, err := json.Marshal(quote)
		w.WriteHeader(http.StatusCreated)
		w.Write(result)
	case http.MethodDelete:
		quoteId, err := strconv.Atoi(r.URL.Path[len("/quotes/"):])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		types.DelQuote(quoteId)
		w.WriteHeader(http.StatusNoContent)
	default:
		msg := "Not implemented method."
		buff := []byte(msg[:])
		w.Write(buff)
	}
}
