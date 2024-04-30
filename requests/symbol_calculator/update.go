package requests_symbol_calculator

import (
	"backend/enums"
	enums_symbol_calculator "backend/enums/symbol_calculator"
)

type UpdateRequest struct {
	Group         string                             `json:"group" validate:"required,alpha,uppercase"`
	Volume        int                                `json:"volume" validate:"number,gte=0"`
	SortColumn    enums_symbol_calculator.SortColumn `json:"sortColumn" validate:"required,symbolCalculatorSortColumn"`
	SortDirection enums.SortDirection                `json:"sortDirection" validate:"required,sortDirection"`
}
