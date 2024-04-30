package models_quote

import (
	"backend/core/services/helper"
	"backend/enums/quote"
)

type QuotePercentModel struct {
	Low  float64 `json:"low"`
	Body float64 `json:"body"`
	High float64 `json:"high"`
}

func GetPercent(direction enums_quote.Direction, priceOpen float64, priceHigh float64, priceLow float64, priceClose float64, fix int) QuotePercentModel {
	if direction == enums_quote.DirectionUp {
		return QuotePercentModel{
			Low:  core_services_helper.GetPercentFromMinMax(priceLow, priceOpen, fix),
			Body: core_services_helper.GetPercentFromMinMax(priceOpen, priceClose, fix),
			High: core_services_helper.GetPercentFromMinMax(priceClose, priceHigh, fix),
		}
	} else {
		return QuotePercentModel{
			High: core_services_helper.GetPercentFromMinMax(priceHigh, priceOpen, fix),
			Body: core_services_helper.GetPercentFromMinMax(priceOpen, priceClose, fix),
			Low:  core_services_helper.GetPercentFromMinMax(priceClose, priceLow, fix),
		}
	}
}
