package services_quote

import (
	core_models "backend/core/models"
	"backend/models/quote"
	variables_calculator "backend/variables/calculator"
	"fmt"
	"sort"
)

func (quoteService *QuoteService) LoadRange(symbol string, quoteRange *models_quote.QuoteRangeModel, progressChan chan *core_models.ProgressChannelModel,
	progress *core_models.ProgressChannelModel) ([]*models_quote.QuoteModel, error) {
	const label = "services.quote.LoadRange"
	var quotes []*models_quote.QuoteModel
	timeSet := make(map[int64]bool)
	timeEnd := quoteRange.TimeTo

	for i := 0; i < quoteRange.Iterations; i++ {
		if variables_calculator.Stop {
			return nil, nil
		}

		result, err := quoteService.Load(symbol, timeEnd, int(quoteRange.Limit))

		if err != nil {
			return nil, fmt.Errorf("%s: %w", label, err)
		}

		for _, quote := range result {
			if !timeSet[quote.TimeOpen] && quote.TimeOpen >= quoteRange.TimeFrom && quote.TimeOpen <= quoteRange.TimeTo {
				timeSet[quote.TimeOpen] = true
				quotes = append(quotes, quote)
			}
		}

		timeEnd -= quoteRange.TimeStep
		progress.Count++
		progressChan <- progress
	}

	sort.Slice(quotes, func(i, j int) bool {
		return quotes[i].TimeOpen < quotes[j].TimeOpen
	})

	return quotes, nil
}
