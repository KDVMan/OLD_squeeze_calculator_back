package services_quote

import (
	"backend/enums"
	"backend/models/quote"
	"fmt"
)

func (quoteService *QuoteService) loadLocal(symbol string, timeEnd int64, limit int) ([]*models_quote.QuoteModel, error) {
	const label = "services.quote.loadLocal"
	var quotes []*models_quote.QuoteModel
	milliseconds := enums.IntervalMilliseconds(enums.Interval1m)

	err := quoteService.storageService.DB.
		Where("symbol = ? AND time_open >= ? AND time_open <= ?", symbol, timeEnd-milliseconds*int64(limit), timeEnd).
		Order("time_open desc").
		Limit(limit).
		Find(&quotes).
		Error

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return quotes, nil
}
