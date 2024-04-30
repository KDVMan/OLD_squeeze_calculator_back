package services_exchange_limit

import (
	models_exchange "backend/models/exchange"
	"fmt"
)

func (exchangeLimitService *ExchangeLimitService) Load() ([]*models_exchange.ExchangeLimitModel, error) {
	const label = "services.exchange-limit.Load"
	var exchangeLimitModel []*models_exchange.ExchangeLimitModel

	if err := exchangeLimitService.storageService.DB.Find(&exchangeLimitModel).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return exchangeLimitModel, nil
}
