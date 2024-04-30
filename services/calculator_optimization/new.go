package services_calculator_optimization

import (
	"backend/models/calculator"
	"backend/models/control_symbol"
)

type CalculatorOptimizationService struct {
	controlCalculatorModel *models_control_calculator.ControlCalculatorModel
	optimizations          map[string]*models_calculator.CalculatorOptimizationModel
}

func New(controlCalculatorModel *models_control_calculator.ControlCalculatorModel) *CalculatorOptimizationService {
	return &CalculatorOptimizationService{
		controlCalculatorModel: controlCalculatorModel,
		optimizations:          make(map[string]*models_calculator.CalculatorOptimizationModel),
	}
}
