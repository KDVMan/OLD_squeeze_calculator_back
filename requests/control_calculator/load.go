package requests_control_calculator

import (
	"backend/enums"
)

type LoadRequest struct {
	Symbol         string               `json:"symbol" validate:"required,alphanum,uppercase"`
	TradeDirection enums.TradeDirection `json:"tradeDirection,omitempty" validate:"omitempty,tradeDirection"`
	Interval       enums.Interval       `json:"interval,omitempty" validate:"omitempty,interval"`
}
