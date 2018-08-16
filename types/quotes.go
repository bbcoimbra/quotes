package types

type Quote struct {
	QuoteId int
	Text    string
	Author  string
}

var (
	quotes         map[int]Quote
	currentQuoteId int
)

func init() {
	quotes = make(map[int]Quote)
}

func AddQuote(q *Quote) {
	currentQuoteId++
	q.QuoteId = currentQuoteId
	quotes[currentQuoteId] = *q
}

func DelQuote(id int) {
	_, ok := quotes[id]
	if ok {
		delete(quotes, id)
	}
}

func ListQuotes() []Quote {
	qs := make([]Quote, 0, len(quotes))
	for _, quote := range quotes {
		qs = append(qs, quote)
	}
	return qs
}

func GetQuote(quoteId int) (quote *Quote) {
	q := quotes[quoteId]
	quote = &q
	return
}
