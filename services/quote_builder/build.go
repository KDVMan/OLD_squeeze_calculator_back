package services_quote_builder

import (
	models_quote "backend/models/quote"
)

func (quoteBuilderService *QuoteBuilderService) Build(quote *models_quote.QuoteModel) *models_quote.QuoteModel {
	if quoteBuilderService.milliseconds == quoteBuilderService.millisecondsSource {
		return quote
	}

	if quoteBuilderService.quote != nil && (quote.TimeOpen/quoteBuilderService.milliseconds) != (quoteBuilderService.quote.TimeOpen/quoteBuilderService.milliseconds) {
		if quoteBuilderService.isLast(quote) {
			quoteBuilderService.quote = nil
		} else {
			newQuote := *quote
			newQuote.IsClosed = false
			quoteBuilderService.quote = &newQuote
		}

		return nil
	}

	if quoteBuilderService.quote != nil {
		quoteBuilderService.quote.TimeClose = quote.TimeClose
		quoteBuilderService.quote.PriceClose = quote.PriceClose
		quoteBuilderService.quote.PriceHigh = max(quote.PriceHigh, quoteBuilderService.quote.PriceHigh)
		quoteBuilderService.quote.PriceLow = min(quote.PriceLow, quoteBuilderService.quote.PriceLow)
		quoteBuilderService.quote.VolumeLeft += quote.VolumeLeft
		quoteBuilderService.quote.VolumeRight += quote.VolumeRight
		quoteBuilderService.quote.Trades += quote.Trades
		quoteBuilderService.quote.VolumeBuyLeft += quote.VolumeBuyLeft
		quoteBuilderService.quote.VolumeBuyRight += quote.VolumeBuyRight
	} else {
		newQuote := *quote
		newQuote.IsClosed = false
		quoteBuilderService.quote = &newQuote
	}

	if quoteBuilderService.isLast(quote) {
		newQuote := *quoteBuilderService.quote
		newQuote.IsClosed = true
		quoteBuilderService.quote = nil

		return &newQuote
	}

	return quoteBuilderService.quote
}

func (quoteBuilderService *QuoteBuilderService) isLast(quote *models_quote.QuoteModel) bool {
	return ((quote.TimeOpen + quoteBuilderService.millisecondsSource) / quoteBuilderService.milliseconds) != (quote.TimeOpen / quoteBuilderService.milliseconds)
}
