package models_init

import (
	"backend/enums"
)

type InitModel struct {
	ID         int              `json:"id" gorm:"primaryKey"`
	Symbol     string           `json:"symbol"`
	Instrument enums.Instrument `json:"instrument"`
}

func (InitModel) TableName() string {
	return "init"
}

func LoadDefault() *InitModel {
	return &InitModel{
		Symbol:     "BTCUSDT",
		Instrument: enums.InstrumentCalculator,
	}
}
