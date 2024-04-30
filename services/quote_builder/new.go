package services_quote_builder

import (
	"backend/enums"
	"backend/models/quote"
)

type QuoteBuilderService struct {
	quote              *models_quote.QuoteModel
	milliseconds       int64
	millisecondsSource int64
}

func New(interval enums.Interval, intervalSource enums.Interval) *QuoteBuilderService {
	return &QuoteBuilderService{
		milliseconds:       enums.IntervalMilliseconds(interval),
		millisecondsSource: enums.IntervalMilliseconds(intervalSource),
	}
}
