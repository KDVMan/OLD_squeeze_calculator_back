package models_symbol

import (
	"backend/enums/symbol"
	"backend/models"
)

type SymbolModel struct {
	models.BaseModel
	Group     string                    `gorm:"not null,index" json:"group"`
	Symbol    string                    `gorm:"unique;not null" json:"symbol"`
	Status    enums_symbol.SymbolStatus `gorm:"not null" json:"status"`
	Limit     SymbolLimitModel          `gorm:"embedded" json:"limit"`
	Statistic SymbolStatisticModel      `gorm:"embedded" json:"statistic"`
}

func (SymbolModel) TableName() string {
	return "symbols"
}
