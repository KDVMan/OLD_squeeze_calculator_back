package services_symbol

import (
	"backend/enums/symbol"
	"backend/models/symbol"
	"backend/requests/symbol"
	"fmt"
)

func (symbolService *SymbolService) Search(request requests_symbol.SearchRequest) ([]models_symbol.SymbolModel, error) {
	const label = "services.symbol.Search"
	var symbols []models_symbol.SymbolModel
	query := fmt.Sprintf("%%%s%%", request.Symbol)

	err := symbolService.storageService.DB.
		Model(&models_symbol.SymbolModel{}).
		Where("UPPER(symbol) LIKE UPPER(?) AND status = ?", query, enums_symbol.SymbolStatusActive).
		Order("symbol").
		Find(&symbols).
		Error

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return symbols, nil
}
