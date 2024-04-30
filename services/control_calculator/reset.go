package services_control_calculator

import (
	models_control_calculator "backend/models/control_symbol"
	requests_control_calculator "backend/requests/control_calculator"
	"fmt"
)

func (controlCalculatorService *ControlCalculatorService) Reset(request requests_control_calculator.ResetRequest) (*models_control_calculator.ControlCalculatorModel, error) {
	const label = "services.control_calculator.Reset"

	controlCalculatorService.storageService.DB.
		Where("symbol = ? AND trade_direction = ? AND interval = ?", request.Symbol, request.TradeDirection, request.Interval).
		Delete(&models_control_calculator.ControlCalculatorModel{})

	controlCalculatorModel, err := controlCalculatorService.Load(requests_control_calculator.LoadRequest(request))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return controlCalculatorModel, nil
}
