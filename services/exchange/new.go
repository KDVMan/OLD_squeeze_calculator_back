package services_exchange

import (
	"backend/services/exchange_limit"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"net/http"
)

type ExchangeService struct {
	exchangeLimitService *services_exchange_limit.ExchangeLimitService
	client               *futures.Client
}

func New(exchangeLimitService *services_exchange_limit.ExchangeLimitService) *ExchangeService {
	client := binance.NewFuturesClient("", "")

	client.HTTPClient = &http.Client{
		Transport: &services_exchange_limit.Transport{Value: http.DefaultTransport},
	}

	return &ExchangeService{
		exchangeLimitService: exchangeLimitService,
		client:               client,
	}
}
