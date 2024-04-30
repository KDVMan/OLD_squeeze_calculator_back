package models_result_calculator

import (
	"backend/enums"
	"backend/models"
)

type ResultCalculatorModel struct {
	models.BaseModel
	Hash                           string                  `gorm:"unique;not null" json:"-"`
	Symbol                         string                  `gorm:"index:index_result_calculator_01,not null" json:"symbol"`
	TradeDirection                 enums.TradeDirection    `gorm:"index:index_result_calculator_01,not null" json:"tradeDirection"`
	Interval                       enums.Interval          `gorm:"index:index_result_calculator_01,not null" json:"interval"`
	ProfitPercent                  float64                 `json:"profitPercent"`
	AverageProfitPercent           float64                 `json:"averageProfitPercent"`
	StandardDeviationProfitPercent float64                 `json:"standardDeviationProfitPercent"`
	Total                          int                     `json:"total"`
	TotalStops                     int                     `json:"totalStops"`
	TotalTakes                     int                     `json:"totalTakes"`
	Coefficient                    float64                 `json:"coefficient"`
	Ratio                          float64                 `json:"ratio"`
	WinRate                        float64                 `json:"winRate"`
	MaxTimeDeal                    int64                   `json:"maxTimeDeal"`
	AverageTimeDeal                int64                   `json:"averageTimeDeal"`
	MaxDrawdown                    float64                 `json:"maxDrawdown"`
	AverageDrawdown                float64                 `json:"averageDrawdown"`
	DrawdownProfitRatio            float64                 `json:"drawdownProfitRatio"`
	Score                          float64                 `json:"score"`
	Param                          *ResultCalculatorParam  `gorm:"embedded" json:"param"`
	Deals                          []*ResultCalculatorDeal `gorm:"foreignKey:ResultCalculatorID;constraint:OnDelete:CASCADE" json:"deals"`
	Filter                         *ResultCalculatorFilter `gorm:"embedded" json:"filter"`
}

func (ResultCalculatorModel) TableName() string {
	return "results_calculators"
}
