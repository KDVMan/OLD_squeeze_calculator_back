package requests_control_calculator

import (
	"backend/enums"
)

type ResetRequest struct {
	Symbol         string               `json:"symbol" validate:"required,alphanum,uppercase"`
	TradeDirection enums.TradeDirection `json:"tradeDirection" validate:"required,tradeDirection"`
	Interval       enums.Interval       `json:"interval" validate:"required,interval"`
}
