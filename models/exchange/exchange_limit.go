package models_exchange

import (
	"backend/enums/exchange"
)

type ExchangeLimitModel struct {
	Type           enums_exchange.RateType     `gorm:"index:index_exchange_limit_01,not null" json:"type"`
	Interval       enums_exchange.RateInterval `gorm:"index:index_exchange_limit_01,not null" json:"interval"`
	IntervalNumber int64                       `gorm:"index:index_exchange_limit_01,not null" json:"intervalNumber"`
	Total          int64                       `gorm:"not null" json:"total"`
	TotalLeft      int64                       `gorm:"not null" json:"totalLeft"`
}

func (ExchangeLimitModel) TableName() string {
	return "exchange_limit"
}
