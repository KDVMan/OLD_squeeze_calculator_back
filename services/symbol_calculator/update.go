package services_symbol_calculator

import (
	"backend/requests/symbol_calculator"
	"fmt"
)

func (symbolCalculatorService *SymbolCalculatorService) Update(request requests_symbol_calculator.UpdateRequest) (*SymbolCalculatorData, error) {
	const label = "services.symbol_calculator.Update"

	symbolCalculatorModel, err := symbolCalculatorService.Load()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	symbolCalculatorModel.Group = request.Group
	symbolCalculatorModel.Volume = request.Volume
	symbolCalculatorModel.SortColumn = request.SortColumn
	symbolCalculatorModel.SortDirection = request.SortDirection

	result := symbolCalculatorService.storageService.DB.Save(symbolCalculatorModel.SymbolCalculatorModel)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", label, result.Error)
	}

	return symbolCalculatorModel, nil
}
