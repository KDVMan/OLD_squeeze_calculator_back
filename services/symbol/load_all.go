package services_symbol

import (
	"backend/enums/symbol"
	"backend/models/symbol"
	"fmt"
)

func (symbolService *SymbolService) LoadAll() (*[]models_symbol.SymbolModel, error) {
	const label = "services.symbol.LoadAll"
	var symbols []models_symbol.SymbolModel

	err := symbolService.storageService.DB.
		Model(&models_symbol.SymbolModel{}).
		Where("status = ?", enums_symbol.SymbolStatusActive).
		Order("symbol").
		Find(&symbols).
		Error

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return &symbols, nil
}
