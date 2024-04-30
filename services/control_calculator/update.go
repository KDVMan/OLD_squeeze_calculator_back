package services_control_calculator

import (
	models_control_calculator "backend/models/control_symbol"
	requests_control_calculator "backend/requests/control_calculator"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (controlCalculatorService *ControlCalculatorService) Update(request requests_control_calculator.StartRequest) (*models_control_calculator.ControlCalculatorModel, error) {
	const label = "services.control_calculator.Update"

	controlCalculatorModel := &models_control_calculator.ControlCalculatorModel{
		Symbol:         request.Symbol,
		TradeDirection: request.TradeDirection,
		Interval:       request.Interval,
	}

	err := controlCalculatorService.storageService.DB.Where(controlCalculatorModel).First(controlCalculatorModel).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	controlCalculatorModel.TimeFrom = request.TimeFrom
	controlCalculatorModel.TimeTo = request.TimeTo
	controlCalculatorModel.OncePerCandle = request.OncePerCandle
	controlCalculatorModel.Bind = request.Bind
	controlCalculatorModel.PercentInFrom = request.PercentInFrom
	controlCalculatorModel.PercentInTo = request.PercentInTo
	controlCalculatorModel.PercentInStep = request.PercentInStep
	controlCalculatorModel.PercentOutFrom = request.PercentOutFrom
	controlCalculatorModel.PercentOutTo = request.PercentOutTo
	controlCalculatorModel.PercentOutStep = request.PercentOutStep
	controlCalculatorModel.StopTime = request.StopTime
	controlCalculatorModel.StopTimeFrom = request.StopTimeFrom
	controlCalculatorModel.StopTimeTo = request.StopTimeTo
	controlCalculatorModel.StopTimeStep = request.StopTimeStep
	controlCalculatorModel.StopPercent = request.StopPercent
	controlCalculatorModel.StopPercentFrom = request.StopPercentFrom
	controlCalculatorModel.StopPercentTo = request.StopPercentTo
	controlCalculatorModel.StopPercentStep = request.StopPercentStep
	controlCalculatorModel.Algorithm = request.Algorithm
	controlCalculatorModel.Iterations = int(request.Iterations)

	if err := controlCalculatorService.storageService.DB.Save(controlCalculatorModel).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return controlCalculatorModel, nil
}
