package services_control_calculator

import (
	"backend/enums"
	models_control_calculator "backend/models/control_symbol"
	requests_control_calculator "backend/requests/control_calculator"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (controlCalculatorService *ControlCalculatorService) Load(request requests_control_calculator.LoadRequest) (*models_control_calculator.ControlCalculatorModel, error) {
	const label = "services.control_calculator.Load"
	var controlCalculatorModel *models_control_calculator.ControlCalculatorModel

	query := controlCalculatorService.storageService.DB.Where("symbol = ?", request.Symbol)

	if request.TradeDirection == "" {
		request.TradeDirection = enums.TradeDirectionLong
	} else {
		query = query.Where("trade_direction = ?", request.TradeDirection)
	}

	if request.Interval == "" {
		request.Interval = enums.Interval1m
	} else {
		query = query.Where("interval = ?", request.Interval)
	}

	if err := query.Order("updated_at desc").First(&controlCalculatorModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models_control_calculator.LoadDefault(request.Symbol, request.TradeDirection, request.Interval), nil
		}

		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return controlCalculatorModel, nil

}
