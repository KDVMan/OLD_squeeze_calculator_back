package services_calculator

import (
	"backend/enums"
	"backend/models/result_calculator"
	"backend/services/quote_builder"
)

func (calculatorService *CalculatorService) Calculate() *models_result_calculator.ResultCalculatorModel {
	quoteBuilderService := services_quote_builder.New(calculatorService.param.Interval, enums.Interval1m)
	var deals []*models_result_calculator.ResultCalculatorDeal
	var deal *models_result_calculator.ResultCalculatorDeal
	nextPriceIn := 0.0

	for i, quote := range calculatorService.quotes {
		quoteBuild := quoteBuilderService.Build(quote)
		currentPriceIn := nextPriceIn

		if quoteBuild.IsClosed {
			nextPriceIn = calculatorService.calculatePriceIn(quoteBuild)
		}

		if deal != nil && deal.TimeOut >= quote.TimeOpen {
			continue
		}

		if calculatorService.param.OncePerCandle && deal != nil && deal.TimeOut >= quoteBuild.TimeOpen {
			continue
		}

		if (calculatorService.param.TradeDirection == enums.TradeDirectionLong && quote.PriceLow < currentPriceIn) ||
			(calculatorService.param.TradeDirection == enums.TradeDirectionShort && currentPriceIn > 0 && quote.PriceHigh >= currentPriceIn) {
			deal = calculatorService.calculateDeal(i, currentPriceIn)
			deals = append(deals, deal)
		}
	}

	return calculatorService.CalculateStatistic(deals)
}
