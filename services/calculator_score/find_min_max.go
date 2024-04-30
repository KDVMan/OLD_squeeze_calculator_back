package services_calculator_score

import (
	"backend/models/result_calculator"
	"math"
)

type MinMaxValues struct {
	Total                          float64
	TotalTakes                     float64
	TotalStops                     float64
	ProfitPercent                  float64
	AverageProfitPercent           float64
	StandardDeviationProfitPercent float64
	WinRate                        float64
	MaxTimeDeal                    int64
	AverageTimeDeal                int64
	Drawdown                       float64
	DrawdownProfitRatio            float64
}

func FindMinMax(results []*models_result_calculator.ResultCalculatorModel) (minValues MinMaxValues, maxValues MinMaxValues) {
	init := func(value float64) MinMaxValues {
		return MinMaxValues{
			Total:                          value,
			TotalTakes:                     value,
			TotalStops:                     value,
			ProfitPercent:                  value,
			AverageProfitPercent:           value,
			StandardDeviationProfitPercent: value,
			WinRate:                        value,
			MaxTimeDeal:                    int64(value),
			AverageTimeDeal:                int64(value),
			Drawdown:                       value,
			DrawdownProfitRatio:            value,
		}
	}

	minValues = init(math.Inf(1))
	maxValues = init(math.Inf(-1))

	for _, result := range results {
		minValues.Total = math.Min(minValues.Total, float64(result.Total))
		maxValues.Total = math.Max(maxValues.Total, float64(result.Total))

		minValues.TotalTakes = math.Min(minValues.TotalTakes, float64(result.TotalTakes))
		maxValues.TotalTakes = math.Max(maxValues.TotalTakes, float64(result.TotalTakes))

		minValues.TotalStops = math.Min(minValues.TotalStops, float64(result.TotalStops))
		maxValues.TotalStops = math.Max(maxValues.TotalStops, float64(result.TotalStops))

		minValues.ProfitPercent = math.Min(minValues.ProfitPercent, result.ProfitPercent)
		maxValues.ProfitPercent = math.Max(maxValues.ProfitPercent, result.ProfitPercent)

		minValues.AverageProfitPercent = math.Min(minValues.AverageProfitPercent, result.AverageProfitPercent)
		maxValues.AverageProfitPercent = math.Max(maxValues.AverageProfitPercent, result.AverageProfitPercent)

		minValues.StandardDeviationProfitPercent = math.Min(minValues.StandardDeviationProfitPercent, result.StandardDeviationProfitPercent)
		maxValues.StandardDeviationProfitPercent = math.Max(maxValues.StandardDeviationProfitPercent, result.StandardDeviationProfitPercent)

		minValues.WinRate = math.Min(minValues.WinRate, result.WinRate)
		maxValues.WinRate = math.Max(maxValues.WinRate, result.WinRate)

		minValues.MaxTimeDeal = int64(math.Min(float64(minValues.MaxTimeDeal), float64(result.MaxTimeDeal)))
		maxValues.MaxTimeDeal = int64(math.Max(float64(maxValues.MaxTimeDeal), float64(result.MaxTimeDeal)))

		minValues.AverageTimeDeal = int64(math.Min(float64(minValues.AverageTimeDeal), float64(result.AverageTimeDeal)))
		maxValues.AverageTimeDeal = int64(math.Max(float64(maxValues.AverageTimeDeal), float64(result.AverageTimeDeal)))

		minValues.Drawdown = math.Min(minValues.Drawdown, result.MaxDrawdown)
		maxValues.Drawdown = math.Max(maxValues.Drawdown, result.MaxDrawdown)

		minValues.DrawdownProfitRatio = math.Min(minValues.DrawdownProfitRatio, result.DrawdownProfitRatio)
		maxValues.DrawdownProfitRatio = math.Max(maxValues.DrawdownProfitRatio, result.DrawdownProfitRatio)
	}

	return
}
