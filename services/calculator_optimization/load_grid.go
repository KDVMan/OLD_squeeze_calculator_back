package services_calculator_optimization

import (
	"backend/core/services/helper"
	models_calculator "backend/models/calculator"
	"fmt"
)

func (calculatorOptimizationService *CalculatorOptimizationService) loadGrid() {
	iteration := 0

	percentInMin, percentInMax, percentInStep, percentInAccuracy := core_services_helper.GetRangeFloatByInt(
		calculatorOptimizationService.controlCalculatorModel.PercentInFrom,
		calculatorOptimizationService.controlCalculatorModel.PercentInTo,
		calculatorOptimizationService.controlCalculatorModel.PercentInStep,
	)

	percentOutMin, percentOutMax, percentOutStep, percentOutAccuracy := core_services_helper.GetRangeFloatByInt(
		calculatorOptimizationService.controlCalculatorModel.PercentOutFrom,
		calculatorOptimizationService.controlCalculatorModel.PercentOutTo,
		calculatorOptimizationService.controlCalculatorModel.PercentOutStep,
	)

	stopPercentMin, stopPercentMax, stopPercentStep, stopPercentAccuracy := core_services_helper.GetRangeFloatByInt(
		calculatorOptimizationService.controlCalculatorModel.StopPercentFrom,
		calculatorOptimizationService.controlCalculatorModel.StopPercentTo,
		calculatorOptimizationService.controlCalculatorModel.StopPercentStep,
	)

	var stopTimeOptions []int64
	var stopPercentOptions []int64

	if calculatorOptimizationService.controlCalculatorModel.StopTime {
		for _, value := range calculatorOptimizationService.generateRange(
			calculatorOptimizationService.controlCalculatorModel.StopTimeFrom,
			calculatorOptimizationService.controlCalculatorModel.StopTimeTo,
			calculatorOptimizationService.controlCalculatorModel.StopTimeStep,
		) {
			stopTimeOptions = append(stopTimeOptions, value)
		}
	}

	if calculatorOptimizationService.controlCalculatorModel.StopPercent {
		for _, value := range calculatorOptimizationService.generateRange(stopPercentMin, stopPercentMax, stopPercentStep) {
			stopPercentOptions = append(stopPercentOptions, value)
		}
	}

	// if calculatorOptimizationService.controlCalculatorModel.StopTime && calculatorOptimizationService.controlCalculatorModel.StopPercent {
	// 	stopTimeOptions = append(stopTimeOptions, -1)
	// 	stopPercentOptions = append(stopPercentOptions, -1)
	// }

	for _, bind := range calculatorOptimizationService.controlCalculatorModel.Bind {
		for percentIn := percentInMin; percentIn <= percentInMax; percentIn += percentInStep {
			for percentOut := percentOutMin; percentOut <= percentOutMax; percentOut += percentOutStep {
				for _, stopTime := range stopTimeOptions {
					adjustedStopTime := stopTime

					if adjustedStopTime > 0 {
						adjustedStopTime = adjustedStopTime * 60 * 1000
					}

					for _, stopPercent := range stopPercentOptions {
						if stopTime == -1 && stopPercent == -1 && calculatorOptimizationService.controlCalculatorModel.StopTime && calculatorOptimizationService.controlCalculatorModel.StopPercent {
							continue
						}

						key := fmt.Sprintf("%s-%d-%d-%d-%d", bind, percentIn, percentOut, stopTime, stopPercent)

						adjustedStopPercent := float64(stopPercent)

						if adjustedStopPercent > 0 {
							adjustedStopPercent = adjustedStopPercent / float64(stopPercentAccuracy)
						}

						if _, exists := calculatorOptimizationService.optimizations[key]; !exists {
							calculatorOptimizationService.optimizations[key] = &models_calculator.CalculatorOptimizationModel{
								Bind:        bind,
								PercentIn:   float64(percentIn) / float64(percentInAccuracy),
								PercentOut:  float64(percentOut) / float64(percentOutAccuracy),
								StopTime:    adjustedStopTime,
								StopPercent: adjustedStopPercent,
							}
						}

						iteration++

						if iteration >= calculatorOptimizationService.controlCalculatorModel.Iterations {
							return
						}
					}
				}
			}
		}
	}
}

func (calculatorOptimizationService *CalculatorOptimizationService) generateRange(from int64, to int64, step int64) []int64 {
	var out []int64

	for value := from; value <= to; value += step {
		out = append(out, value)
	}

	return out
}
