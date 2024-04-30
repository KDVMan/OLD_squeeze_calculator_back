package services_exchange

import (
	"backend/services/exchange_limit"
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
)

func (exchangeService *ExchangeService) Kline(symbol string, interval string, timeEnd int64, limit int) ([]*futures.Kline, error) {
	const label = "services.exchange.Kline"

	result, err := exchangeService.client.NewKlinesService().
		Symbol(symbol).
		Interval(interval).
		EndTime(timeEnd).
		Limit(limit).
		Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	err = exchangeService.exchangeLimitService.Update(services_exchange_limit.GetLimits())

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return result, nil
}
