package requests_init

import (
	"backend/enums"
)

type UpdateRequest struct {
	Symbol     string           `json:"symbol" validate:"required,alphanum,uppercase"`
	Instrument enums.Instrument `json:"instrument" validate:"required,instrument"`
}
