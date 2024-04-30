package services_calculator

import (
	"backend/enums"
	"backend/enums/calculator"
	"backend/models/result_calculator"
)

func (calculatorService *CalculatorService) calculateDeal(index int, priceIn float64) *models_result_calculator.ResultCalculatorDeal {
	deal := &models_result_calculator.ResultCalculatorDeal{
		TimeIn:        calculatorService.quotes[index].TimeOpen,
		TimeOut:       0,
		PriceIn:       priceIn,
		PriceOut:      0,
		ProfitPercent: 0,
	}

	priceOut := calculatorService.calculatePriceOut(priceIn)
	priceStop := calculatorService.calculatePriceStop(priceIn)

	check := calculatorService.checkDeal(calculatorService.quotes[index], deal, priceOut, priceStop, enums_calculator.DealPriceClose)
	index++

	for check == false && index < len(calculatorService.quotes) {
		check = calculatorService.checkDeal(calculatorService.quotes[index], deal, priceOut, priceStop, calculatorService.getBindDeal())
		index++
	}

	if check == false {
		deal.TimeOut = calculatorService.quotes[len(calculatorService.quotes)-1].TimeClose
		deal.PriceOut = calculatorService.quotes[len(calculatorService.quotes)-1].PriceClose
	}

	costIn := deal.PriceIn * (1 + calculatorService.commission/100)
	costOut := deal.PriceOut * (1 - calculatorService.commission/100)

	if calculatorService.param.TradeDirection == enums.TradeDirectionLong {
		deal.ProfitPercent = (costOut/costIn - 1) * 100
	} else if calculatorService.param.TradeDirection == enums.TradeDirectionShort {
		deal.ProfitPercent = (costIn/costOut - 1) * 100
	}

	return deal
}
