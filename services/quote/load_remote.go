package services_quote

import (
	"backend/enums"
	"backend/models/quote"
	"fmt"
	"gorm.io/gorm/clause"
	"time"
)

func (quoteService *QuoteService) loadRemote(symbol string, timeEnd int64, limit int) ([]*models_quote.QuoteModel, error) {
	const label = "services.quote.loadRemote"
	var quotes []*models_quote.QuoteModel
	milliseconds := enums.IntervalMilliseconds(enums.Interval1m)

	klines, err := quoteService.exchangeService.Kline(symbol, string(enums.Interval1m), timeEnd, limit)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	tx := quoteService.storageService.DB.Begin()

	if tx.Error != nil {
		return nil, fmt.Errorf("%s: %w", label, tx.Error)
	}

	for _, kline := range klines {
		quote := models_quote.KlineToQuote(symbol, kline)

		if quote.TimeClose <= time.Now().UnixMilli() { // пишем в базу только закрытые свечи
			checkTime := (quote.TimeClose - quote.TimeOpen) + 1

			if checkTime < milliseconds {
				quote.TimeClose = quote.TimeOpen + milliseconds - 1 // битые данные (бинанс так иногда отдает)
			}

			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&quote).Error

			if err != nil {
				tx.Rollback()
				return nil, fmt.Errorf("%s: %w", label, err)
			}
		}

		quotes = append(quotes, quote)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return quotes, nil
}
