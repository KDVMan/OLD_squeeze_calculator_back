package requests_result_calculator

import (
	"backend/enums"
	enums_result_calculator "backend/enums/result_calculator"
)

type LoadRequest struct {
	Symbol         string                             `json:"symbol" validate:"required,alphanum,uppercase"`
	TradeDirection enums.TradeDirection               `json:"tradeDirection" validate:"required,tradeDirection"`
	Interval       enums.Interval                     `json:"interval" validate:"required,interval"`
	Limit          int                                `json:"limit" validate:"required,gt=0"`
	SortColumn     enums_result_calculator.SortColumn `json:"sortColumn" validate:"required,resultCalculatorSortColumn"`
	SortDirection  enums.SortDirection                `json:"sortDirection" validate:"required,sortDirection"`
}
