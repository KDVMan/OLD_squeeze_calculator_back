package services_calculator

import (
	core_services_helper "backend/core/services/helper"
	"backend/models/result_calculator"
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func (calculatorService *CalculatorService) CalculateStatistic2(deals []*models_result_calculator.ResultCalculatorDeal) *models_result_calculator.ResultCalculatorModel {
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

	var totalTimeDeal int64
	var maxProfit float64
	var stops float64
	var takes float64
	var drawdowns []float64
	var profitPercents []float64

	for _, deal := range deals {
		result.ProfitPercent += deal.ProfitPercent
		profitPercents = append(profitPercents, deal.ProfitPercent)
		timeDeal := deal.TimeOut - deal.TimeIn
		totalTimeDeal += timeDeal

		if timeDeal > result.MaxTimeDeal {
			result.MaxTimeDeal = timeDeal
		}

		if deal.IsStopTime || deal.IsStopPercent {
			result.TotalStops++
		} else {
			result.TotalTakes++
		}

		if deal.ProfitPercent < 0 {
			stops += deal.ProfitPercent
		} else {
			takes += deal.ProfitPercent
		}

		if result.ProfitPercent > maxProfit {
			maxProfit = result.ProfitPercent
		} else {
			drawdown := maxProfit - result.ProfitPercent
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
