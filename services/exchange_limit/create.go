package services_exchange_limit

import (
	"backend/enums/exchange"
	"backend/models/exchange"
	"fmt"
	"github.com/adshao/go-binance/v2/futures"
)

func (exchangeLimitService *ExchangeLimitService) Create(limits []futures.RateLimit) error {
	const label = "services.exchange-limit.Create"

	for _, limit := range limits {
		exchangeLimitModel := models_exchange.ExchangeLimitModel{
			Type:           convertType(limit.RateLimitType),
			Interval:       convertInterval(limit.Interval),
			IntervalNumber: limit.IntervalNum,
			Total:          limit.Limit,
			TotalLeft:      limit.Limit,
		}

		err := exchangeLimitService.storageService.DB.
			Where("type = ? AND interval = ? AND interval_number = ?", exchangeLimitModel.Type, exchangeLimitModel.Interval, exchangeLimitModel.IntervalNumber).
			Assign(exchangeLimitModel).
			FirstOrCreate(&exchangeLimitModel).
			Error

		if err != nil {
			return fmt.Errorf("%s: %w", label, err)
		}
	}

	return nil
}

func convertType(input string) enums_exchange.RateType {
	switch input {
	case "REQUEST_WEIGHT":
		return enums_exchange.RateTypeWeight
	case "ORDERS":
		return enums_exchange.RateTypeOrder
	default:
		return enums_exchange.RateTypeUnknown
	}
}

func convertInterval(input string) enums_exchange.RateInterval {
	switch input {
	case "SECOND":
		return enums_exchange.RateIntervalSecond
	case "MINUTE":
		return enums_exchange.RateIntervalMinute
	case "HOUR":
		return enums_exchange.RateIntervalHour
	case "DAY":
		return enums_exchange.RateIntervalDay
	default:
		return enums_exchange.RateIntervalUnknown
	}
}
