package services_symbol

import (
	"backend/enums/symbol"
	"backend/models/symbol"
	"fmt"
)

func (symbolService *SymbolService) Load(symbol string, status enums_symbol.SymbolStatus) (*models_symbol.SymbolModel, error) {
	const label = "services.symbol.Load"
	var symbolModel models_symbol.SymbolModel

	if err := symbolService.storageService.DB.Where("symbol = ? AND status = ?", symbol, status).First(&symbolModel).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return &symbolModel, nil
}
