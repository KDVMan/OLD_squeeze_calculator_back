package models_symbol_calculator

import (
	"backend/enums"
	"backend/enums/symbol_calculator"
)

type SymbolCalculatorModel struct {
	ID            int                                `json:"id" gorm:"primaryKey"`
	Group         string                             `json:"group"`
	Volume        int                                `json:"volume"`
	SortColumn    enums_symbol_calculator.SortColumn `json:"sortColumn"`
	SortDirection enums.SortDirection                `json:"sortDirection"`
}

func (SymbolCalculatorModel) TableName() string {
	return "symbol_calculator"
}

func LoadDefault() *SymbolCalculatorModel {
	return &SymbolCalculatorModel{
		Group:         "USDT",
		Volume:        0,
		SortColumn:    enums_symbol_calculator.SortColumnVolume,
		SortDirection: enums.SortDirectionDesc,
	}
}
