package services_symbol_calculator

import (
	"backend/models/symbol"
	"backend/models/symbol_calculator"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type SymbolCalculatorData struct {
	*models_symbol_calculator.SymbolCalculatorModel
	Groups []string `json:"groups"`
}

func (symbolCalculatorService *SymbolCalculatorService) Load() (*SymbolCalculatorData, error) {
	const label = "services.symbol_calculator.Load"
	var symbolCalculatorModel *models_symbol_calculator.SymbolCalculatorModel

	if err := symbolCalculatorService.storageService.DB.First(&symbolCalculatorModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			symbolCalculatorModel = models_symbol_calculator.LoadDefault()

			if err := symbolCalculatorService.storageService.DB.Create(symbolCalculatorModel).Error; err != nil {
				return nil, fmt.Errorf("%s: %w", label, err)
			}

			return &SymbolCalculatorData{
				SymbolCalculatorModel: symbolCalculatorModel,
				Groups:                []string{},
			}, nil
		}

		return nil, fmt.Errorf("%s: %w", label, err)
	}

	groups, err := symbolCalculatorService.LoadGroups()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return &SymbolCalculatorData{
		SymbolCalculatorModel: symbolCalculatorModel,
		Groups:                groups,
	}, nil
}

func (symbolCalculatorService *SymbolCalculatorService) LoadGroups() ([]string, error) {
	var groups []string

	if err := symbolCalculatorService.storageService.DB.
		Model(&models_symbol.SymbolModel{}).
		Distinct("`group`").
		Order("`group` ASC").
		Pluck("`group`", &groups).Error; err != nil {
		return nil, err
	}

	return groups, nil
}
