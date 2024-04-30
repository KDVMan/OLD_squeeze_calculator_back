package services_exchange

import (
	"backend/services/exchange_limit"
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
)

func (exchangeService *ExchangeService) ExchangeInfo() ([]futures.Symbol, error) {
	const label = "services.exchange.ExchangeInfo"

	result, err := exchangeService.client.NewExchangeInfoService().Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	err = exchangeService.exchangeLimitService.Create(result.RateLimits)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	err = exchangeService.exchangeLimitService.Update(services_exchange_limit.GetLimits())

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return result.Symbols, nil
}
