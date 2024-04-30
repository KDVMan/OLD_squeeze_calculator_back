package services_calculator

import (
	core_services_helper "backend/core/services/helper"
	"backend/enums"
	models_quote "backend/models/quote"
)

func (calculatorService *CalculatorService) calculatePriceIn(quote *models_quote.QuoteModel) float64 {
	priceBind := calculatorService.calculatePriceBind(calculatorService.param.Bind, quote)
	priceIn := priceBind * calculatorService.priceInFactor

	return core_services_helper.Floor(priceIn, calculatorService.tickSizeFactor)
}

func (calculatorService *CalculatorService) calculatePriceOut(priceIn float64) float64 {
	return core_services_helper.Floor(priceIn*calculatorService.priceOutFactor, calculatorService.tickSizeFactor)
}

func (calculatorService *CalculatorService) calculatePriceStop(priceIn float64) float64 {
	if calculatorService.priceStopFactor <= 0 {
		return 0
	}

	return core_services_helper.Floor(priceIn*calculatorService.priceStopFactor, calculatorService.tickSizeFactor)
}

func (calculatorService *CalculatorService) calculatePriceBind(bind enums.Bind, quote *models_quote.QuoteModel) float64 {
	switch bind {
	case enums.BindLow:
		return quote.PriceLow
	case enums.BindHigh:
		return quote.PriceHigh
	case enums.BindOpen:
		return quote.PriceOpen
	case enums.BindClose:
		return quote.PriceClose
	case enums.BindMhl:
		return (quote.PriceHigh + quote.PriceLow) / 2
	case enums.BindMoc:
		return (quote.PriceOpen + quote.PriceClose) / 2
	default:
		return 0
	}
}
