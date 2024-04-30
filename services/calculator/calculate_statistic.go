package services_calculator

import (
	"backend/core/services/helper"
	"backend/models/result_calculator"
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func (calculatorService *CalculatorService) CalculateStatistic(deals []*models_result_calculator.ResultCalculatorDeal) *models_result_calculator.ResultCalculatorModel {
	capital := 100.0
	maxCapital := capital
	var drawdowns []float64
	totalTimeDeal := int64(0)
	totalProfitPercent := 0.0
	var profitPercents []float64
	stops := 0.0
	takes := 0.0

	paramJSON, _ := json.Marshal(calculatorService.param)
	hash := fmt.Sprintf("%x", md5.Sum(paramJSON))

	result := &models_result_calculator.ResultCalculatorModel{
		Hash:                           hash,
		Symbol:                         calculatorService.param.Symbol,
		TradeDirection:                 calculatorService.param.TradeDirection,
		Interval:                       calculatorService.param.Interval,
		ProfitPercent:                  0,
		AverageProfitPercent:           0,
		StandardDeviationProfitPercent: 0,
		Total:                          len(deals),
		TotalStops:                     0,
		TotalTakes:                     0,
		Coefficient:                    0,
		Ratio:                          calculatorService.param.PercentOut / calculatorService.param.PercentIn,
		WinRate:                        0,
		MaxTimeDeal:                    0,
		AverageTimeDeal:                0,
		MaxDrawdown:                    0,
		AverageDrawdown:                0,
		DrawdownProfitRatio:            0,
		Score:                          0,
		Param:                          calculatorService.param,
		Deals:                          deals,
	}

	for _, deal := range deals {
		dealProfit := deal.ProfitPercent / 100
		capital = capital * (1 + dealProfit)
		totalProfitPercent += deal.ProfitPercent
		profitPercents = append(profitPercents, deal.ProfitPercent)

		timeDeal := deal.TimeOut - deal.TimeIn
		totalTimeDeal += timeDeal

		if timeDeal > result.MaxTimeDeal {
			result.MaxTimeDeal = timeDeal
		}

		if deal.IsStopTime || deal.IsStopPercent {
			result.TotalStops += 1
		} else {
			result.TotalTakes += 1
		}

		if deal.ProfitPercent < 0 {
			stops += deal.ProfitPercent
		} else {
			takes += deal.ProfitPercent
		}

		if capital > maxCapital {
			maxCapital = capital
		} else {
			drawdown := (maxCapital - capital) / maxCapital * 100
			drawdowns = append(drawdowns, drawdown)

			if drawdown > result.MaxDrawdown {
				result.MaxDrawdown = drawdown
			}
		}
	}

	if stops != 0 {
		result.Coefficient = -takes / stops
	}

	if len(deals) > 0 {
		result.ProfitPercent = ((capital - 100) / 100) * 100
		result.AverageProfitPercent = result.ProfitPercent / float64(len(deals))
		result.StandardDeviationProfitPercent = core_services_helper.CalculateStandardDeviation(profitPercents, result.AverageProfitPercent)
		result.AverageTimeDeal = totalTimeDeal / int64(len(deals))
		result.AverageDrawdown = core_services_helper.CalculateAverage(drawdowns)
		result.WinRate = float64(result.TotalTakes) / float64(len(deals))

		if result.MaxDrawdown > 0 {
			result.DrawdownProfitRatio = result.ProfitPercent / result.MaxDrawdown
		}
	}

	return result
}
