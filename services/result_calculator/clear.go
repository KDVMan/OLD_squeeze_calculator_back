package services_result_calculator

import (
	"backend/enums"
	"backend/models/result_calculator"
	"fmt"
)

func (resultCalculatorService *ResultCalculatorService) Clear(symbol string, tradeDirection enums.TradeDirection, interval enums.Interval) error {
	const label = "services.result_calculator.Clear"

	err := resultCalculatorService.storageService.DB.
		Where("symbol = ? AND trade_direction = ? AND interval = ?", symbol, tradeDirection, interval).
		Delete(&models_result_calculator.ResultCalculatorModel{}).
		Error

	if err != nil {
		return fmt.Errorf("%s: %w", label, err)
	}

	return nil
}
