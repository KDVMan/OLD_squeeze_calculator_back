package services_calculator

import (
	"backend/enums"
	"backend/models/quote"
	"backend/models/result_calculator"
)

type CalculatorService struct {
	param           *models_result_calculator.ResultCalculatorParam
	quotes          []*models_quote.QuoteModel
	commission      float64
	priceInFactor   float64
	priceOutFactor  float64
	priceStopFactor float64
	tickSizeFactor  int
}

func New(param *models_result_calculator.ResultCalculatorParam, quotes []*models_quote.QuoteModel, tickSize float64, commission float64) *CalculatorService {
	calculatorService := &CalculatorService{
		param:      param,
		quotes:     quotes,
		commission: commission,
	}

	if param.TradeDirection == enums.TradeDirectionLong {
		calculatorService.priceInFactor = (100 - param.PercentIn) / 100
		calculatorService.priceOutFactor = (100 + param.PercentOut) / 100

		if param.StopPercent > 0 {
			calculatorService.priceStopFactor = (100 - param.StopPercent) / 100
		}
	} else {
		calculatorService.priceInFactor = (100 + param.PercentIn) / 100
		calculatorService.priceOutFactor = (100 - param.PercentOut) / 100

		if param.StopPercent > 0 {
			calculatorService.priceStopFactor = (100 + param.StopPercent) / 100
		}
	}

	for tickSize < 1 {
		tickSize *= 10
		calculatorService.tickSizeFactor++
	}

	return calculatorService
}
