package services_exchange_limit

import (
	"backend/core/models"
	"backend/enums"
	"backend/enums/exchange"
	"backend/models/exchange"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func (exchangeLimitService *ExchangeLimitService) Update(limits map[string]int) error {
	const label = "services.exchange_limit.Update"

	for key, used := range limits {
		data := map[string]interface{}{
			"total_left": gorm.Expr("total - ?", used),
		}

		err := exchangeLimitService.storageService.DB.
			Model(&models_exchange.ExchangeLimitModel{}).
			Where("type = ? AND interval = ?", getType(key), getInterval(key)).
			Updates(data).
			Error

		if err != nil {
			return fmt.Errorf("%s: %w", label, err)
		}
	}

	exchangeLimitModel, err := exchangeLimitService.Load()

	if err != nil {
		return fmt.Errorf("%s: %w", label, err)
	}

	broadcastModel := core_models.BroadcastChannelModel{
		Event: enums.WebsocketEventExchangeLimits,
		Data:  exchangeLimitModel,
	}

	exchangeLimitService.broadcastChan <- &broadcastModel

	return nil
}

func getType(key string) enums_exchange.RateType {
	switch key {
	case "x-mbx-used-weight", "x-mbx-used-weight-1m":
		return enums_exchange.RateTypeWeight
	case "x-mbx-order-count-1s", "x-mbx-order-count-1m", "x-mbx-order-count-1h", "x-mbx-order-count-1d":
		return enums_exchange.RateTypeOrder
	default:
		return enums_exchange.RateTypeUnknown
	}
}

func getInterval(key string) enums_exchange.RateInterval {
	if strings.Contains(key, "1s") {
		return enums_exchange.RateIntervalSecond
	} else if strings.Contains(key, "1m") {
		return enums_exchange.RateIntervalMinute
	} else if strings.Contains(key, "1h") {
		return enums_exchange.RateIntervalHour
	} else if strings.Contains(key, "1d") {
		return enums_exchange.RateIntervalDay
	}

	return enums_exchange.RateIntervalUnknown
}
