package models_result_calculator

import "backend/models"

type ResultCalculatorDeal struct {
	models.BaseModel
	ResultCalculatorID uint    `gorm:"index:index_result_calculator_deal_01,not null" json:"-"`
	TimeIn             int64   `json:"timeIn"`
	TimeOut            int64   `json:"timeOut"`
	PriceIn            float64 `json:"priceIn"`
	PriceOut           float64 `json:"priceOut"`
	ProfitPercent      float64 `json:"profitPercent"`
	IsStopTime         bool    `json:"isStopTime"`
	IsStopPercent      bool    `json:"isStopPercent"`
}

func (ResultCalculatorDeal) TableName() string {
	return "results_calculators_deals"
}
