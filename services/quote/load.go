package services_quote

import (
	"backend/models/quote"
	"fmt"
)

func (quoteService *QuoteService) Load(symbol string, timeEnd int64, limit int) ([]*models_quote.QuoteModel, error) {
	const label = "services.quote.Load"

	quotes, err := quoteService.loadLocal(symbol, timeEnd, limit)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	if len(quotes) < limit {
		remoteQuotes, err := quoteService.loadRemote(symbol, timeEnd, limit)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", label, err)
		}

		quotes = append(quotes, remoteQuotes...)
	}

	return quotes, nil
}
