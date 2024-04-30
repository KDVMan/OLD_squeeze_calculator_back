package services_calculator_optimization

import (
	"backend/enums"
	"backend/models/calculator"
)

func (calculatorOptimizationService *CalculatorOptimizationService) Load() []*models_calculator.CalculatorOptimizationModel {
	if calculatorOptimizationService.controlCalculatorModel.Algorithm == enums.AlgorithmRandom {
		calculatorOptimizationService.loadRandom()
	} else if calculatorOptimizationService.controlCalculatorModel.Algorithm == enums.AlgorithmGrid {
		calculatorOptimizationService.loadGrid()
	}

	optimizations := make([]*models_calculator.CalculatorOptimizationModel, 0, len(calculatorOptimizationService.optimizations))

	for _, optimization := range calculatorOptimizationService.optimizations {
		optimizations = append(optimizations, optimization)
	}

	return optimizations
}
