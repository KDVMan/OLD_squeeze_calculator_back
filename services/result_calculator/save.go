package services_result_calculator

import (
	"backend/models/result_calculator"
	"fmt"
)

func (resultCalculatorService *ResultCalculatorService) Save(results []*models_result_calculator.ResultCalculatorModel) error {
	const label = "services.result_calculator.Save"
	const batchSize = 50

	if err := resultCalculatorService.storageService.DB.CreateInBatches(results, batchSize).Error; err != nil {
		return fmt.Errorf("%s: %w", label, err)
	}

	return nil
}
