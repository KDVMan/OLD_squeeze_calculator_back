package models_result_calculator

import "backend/enums"

type ResultCalculatorParam struct {
	Symbol         string               `gorm:"column:param_symbol;-" json:"symbol"`
	TradeDirection enums.TradeDirection `gorm:"column:param_trade_direction;-" json:"tradeDirection"`
	Interval       enums.Interval       `gorm:"column:param_interval;-" json:"interval"`
	Bind           enums.Bind           `gorm:"column:param_bind" json:"bind"`
	PercentIn      float64              `gorm:"column:param_percent_in" json:"percentIn"`
	PercentOut     float64              `gorm:"column:param_percent_out" json:"percentOut"`
	StopTime       int64                `gorm:"column:param_stop_time" json:"stopTime"`
	StopPercent    float64              `gorm:"column:param_stop_percent" json:"stopPercent"`
	OncePerCandle  bool                 `gorm:"column:param_once_per_candle" json:"oncePerCandle"`
}
