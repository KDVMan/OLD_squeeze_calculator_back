package services_calculator_score

import (
	"backend/models/result_calculator"
)

func Score(result *models_result_calculator.ResultCalculatorModel, minValues MinMaxValues, maxValues MinMaxValues) float64 {
	normalize := func(value float64, min float64, max float64) float64 {
		if max > min {
			return (value - min) / (max - min)
		}

		return 0
	}

	score := 0.0

	// положительные
	// score += normalize(float64(result.TotalTakes), minValues.TotalTakes, maxValues.TotalTakes) * 0.04
	score += normalize(result.ProfitPercent, minValues.ProfitPercent, maxValues.ProfitPercent) * 1.6
	score += normalize(result.DrawdownProfitRatio, minValues.DrawdownProfitRatio, maxValues.DrawdownProfitRatio) * 0.02
	score += normalize(result.WinRate, minValues.WinRate, maxValues.WinRate) * 0.4

	// стабильность
	score += normalize(result.StandardDeviationProfitPercent, minValues.StandardDeviationProfitPercent, maxValues.StandardDeviationProfitPercent) * -1.0

	// отрицательные
	score -= normalize(result.MaxDrawdown, minValues.Drawdown, maxValues.Drawdown) * 0.9
	// score -= normalize(float64(result.AverageTimeDeal), float64(minValues.MaxTimeDeal), float64(maxValues.MaxTimeDeal)) * 0.4
	// score -= normalize(float64(result.AverageTimeDeal), float64(minValues.AverageTimeDeal), float64(maxValues.AverageTimeDeal)) * 0.2

	return score
}
