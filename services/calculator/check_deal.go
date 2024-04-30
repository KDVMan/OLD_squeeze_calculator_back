package services_calculator

import (
	"backend/enums"
	enums_calculator "backend/enums/calculator"
	models_quote "backend/models/quote"
	"backend/models/result_calculator"
	"log"
)

func (calculatorService *CalculatorService) checkDeal(quote *models_quote.QuoteModel, deal *models_result_calculator.ResultCalculatorDeal, priceOut float64, priceStop float64,
	bind enums_calculator.Deal) bool {
	var currentPrice float64

	switch bind {
	case enums_calculator.DealPriceLow:
		currentPrice = quote.PriceLow
	case enums_calculator.DealPriceHigh:
		currentPrice = quote.PriceHigh
	case enums_calculator.DealPriceClose:
		currentPrice = quote.PriceClose
	default:
		log.Println("Invalid bind value", bind)
		return false
	}

	if priceStop > 0 {
		if (calculatorService.param.TradeDirection == enums.TradeDirectionLong && quote.PriceLow <= priceStop) ||
			(calculatorService.param.TradeDirection == enums.TradeDirectionShort && quote.PriceHigh >= priceStop) {
			deal.IsStopPercent = true
			deal.TimeOut = quote.TimeClose
			deal.PriceOut = priceStop

			return true
		}
	}

	if calculatorService.param.StopTime > 0 && (quote.TimeClose-deal.TimeIn > calculatorService.param.StopTime) {
		deal.IsStopTime = true
		deal.TimeOut = quote.TimeClose

		if calculatorService.param.TradeDirection == enums.TradeDirectionLong {
			deal.PriceOut = quote.PriceLow
		} else {
			deal.PriceOut = quote.PriceHigh
		}

		return true
	}

	if (calculatorService.param.TradeDirection == enums.TradeDirectionLong && currentPrice > priceOut) ||
		(calculatorService.param.TradeDirection == enums.TradeDirectionShort && currentPrice <= priceOut) {
		deal.TimeOut = quote.TimeClose
		deal.PriceOut = priceOut

		return true
	}

	return false
}

func (calculatorService *CalculatorService) getBindDeal() enums_calculator.Deal {
	if calculatorService.param.TradeDirection == enums.TradeDirectionLong {
		return enums_calculator.DealPriceHigh
	} else {
		return enums_calculator.DealPriceLow
	}
}
