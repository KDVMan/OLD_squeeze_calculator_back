package services_result_calculator

import (
	"backend/core/services/helper"
	"backend/models/result_calculator"
	"backend/requests/result_calculator"
	"fmt"
	"log"
)

func (resultCalculatorService *ResultCalculatorService) Load(request requests_result_calculator.LoadRequest) ([]*models_result_calculator.ResultCalculatorModel, error) {
	const label = "services.control_calculator.Load"
	var results []*models_result_calculator.ResultCalculatorModel

	query := resultCalculatorService.storageService.DB.
		Where("symbol = ? AND trade_direction = ? AND interval = ?", request.Symbol, request.TradeDirection, request.Interval).
		Order(fmt.Sprintf("%s %s", core_services_helper.ToSnakeCase(string(request.SortColumn)), request.SortDirection))

	log.Println("query", fmt.Sprintf("%s %s", request.SortColumn, request.SortDirection))

	if err := query.Limit(request.Limit).Find(&results).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return results, nil
}
