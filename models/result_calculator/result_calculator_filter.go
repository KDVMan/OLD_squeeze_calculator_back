package models_result_calculator

import (
	"backend/enums"
	enums_result_calculator "backend/enums/result_calculator"
)

type ResultCalculatorFilter struct {
	ProfitPercent float64                            `gorm:"column:filter_profit_percent" json:"profitPercent"`
	Total         float64                            `gorm:"column:filter_total" json:"total"`
	Coefficient   float64                            `gorm:"column:filter_coefficient" json:"coefficient"`
	Ratio         float64                            `gorm:"column:filter_ratio" json:"ratio"`
	WinRate       float64                            `gorm:"column:filter_win_rate" json:"winRate"`
	Score         float64                            `gorm:"column:filter_score" json:"score"`
	SortColumn    enums_result_calculator.SortColumn `gorm:"column:filter_sort_column" json:"sortColumn"`
	SortDirection enums.SortDirection                `gorm:"column:filter_sort_direction" json:"sortDirection"`
}

func LoadDefault() *ResultCalculatorFilter {
	return &ResultCalculatorFilter{
		ProfitPercent: 0,
		Total:         0,
		Coefficient:   0,
		Ratio:         0,
		WinRate:       0,
		Score:         0,
		SortColumn:    enums_result_calculator.SortColumnScore,
		SortDirection: enums.SortDirectionDesc,
	}
}
