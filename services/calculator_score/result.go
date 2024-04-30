package services_calculator_score

import "backend/models/result_calculator"

func Results(results []*models_result_calculator.ResultCalculatorModel) []*models_result_calculator.ResultCalculatorModel {
	minValues, maxValues := FindMinMax(results)
	filteredResults := make([]*models_result_calculator.ResultCalculatorModel, 0)

	for _, result := range results {
		if result.ProfitPercent >= 5 {
			result.Deals = nil
			result.Score = Score(result, minValues, maxValues)
			filteredResults = append(filteredResults, result)
		}
	}

	return filteredResults
}
