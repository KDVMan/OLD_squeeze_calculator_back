package services_calculator_optimization

import (
	"backend/core/services/helper"
	"backend/enums"
	"backend/models/calculator"
	"fmt"
)

func (calculatorOptimizationService *CalculatorOptimizationService) loadRandom() {
	for iteration := 0; iteration < calculatorOptimizationService.controlCalculatorModel.Iterations; iteration++ {
		var stopTime int64 = -1
		var stopPercent float64 = -1

		bind := enums.BindRandom(calculatorOptimizationService.controlCalculatorModel.Bind)

		percentIn := core_services_helper.GetRandomFloatByInt(
			calculatorOptimizationService.controlCalculatorModel.PercentInFrom,
			calculatorOptimizationService.controlCalculatorModel.PercentInTo,
			calculatorOptimizationService.controlCalculatorModel.PercentInStep,
		)

		percentOut := core_services_helper.GetRandomFloatByInt(
			calculatorOptimizationService.controlCalculatorModel.PercentOutFrom,
			calculatorOptimizationService.controlCalculatorModel.PercentOutTo,
			calculatorOptimizationService.controlCalculatorModel.PercentOutStep,
		)

		if calculatorOptimizationService.controlCalculatorModel.StopTime && calculatorOptimizationService.controlCalculatorModel.StopPercent {
			// decision := rand.Float64()
			//
			// if decision < 0.33 {
			// 	stopTime = core_services_helper.GetRandomInt(
			// 		calculatorOptimizationService.controlCalculatorModel.StopTimeFrom,
			// 		calculatorOptimizationService.controlCalculatorModel.StopTimeTo,
			// 		calculatorOptimizationService.controlCalculatorModel.StopTimeStep,
			// 	)
			// } else if decision < 0.66 {
			// 	stopPercent = core_services_helper.GetRandomFloatByInt(
			// 		calculatorOptimizationService.controlCalculatorModel.StopPercentFrom,
			// 		calculatorOptimizationService.controlCalculatorModel.StopPercentTo,
			// 		calculatorOptimizationService.controlCalculatorModel.StopPercentStep,
			// 	)
			// } else {
			stopTime = core_services_helper.GetRandomInt(
				calculatorOptimizationService.controlCalculatorModel.StopTimeFrom,
				calculatorOptimizationService.controlCalculatorModel.StopTimeTo,
				calculatorOptimizationService.controlCalculatorModel.StopTimeStep,
			)

			stopPercent = core_services_helper.GetRandomFloatByInt(
				calculatorOptimizationService.controlCalculatorModel.StopPercentFrom,
				calculatorOptimizationService.controlCalculatorModel.StopPercentTo,
				calculatorOptimizationService.controlCalculatorModel.StopPercentStep,
			)
			// }
		} else if calculatorOptimizationService.controlCalculatorModel.StopTime {
			stopTime = core_services_helper.GetRandomInt(
				calculatorOptimizationService.controlCalculatorModel.StopTimeFrom,
				calculatorOptimizationService.controlCalculatorModel.StopTimeTo,
				calculatorOptimizationService.controlCalculatorModel.StopTimeStep,
			)
		} else if calculatorOptimizationService.controlCalculatorModel.StopPercent {
			stopPercent = core_services_helper.GetRandomFloatByInt(
				calculatorOptimizationService.controlCalculatorModel.StopPercentFrom,
				calculatorOptimizationService.controlCalculatorModel.StopPercentTo,
				calculatorOptimizationService.controlCalculatorModel.StopPercentStep,
			)
		}

		adjustedStopTime := stopTime

		if adjustedStopTime > 0 {
			adjustedStopTime = adjustedStopTime * 60 * 1000
		}

		key := fmt.Sprintf("%s-%g-%g-%d-%g", bind, percentIn, percentOut, stopTime, stopPercent)

		if _, exists := calculatorOptimizationService.optimizations[key]; !exists {
			calculatorOptimizationService.optimizations[key] = &models_calculator.CalculatorOptimizationModel{
				Bind:        bind,
				PercentIn:   percentIn,
				PercentOut:  percentOut,
				StopTime:    adjustedStopTime,
				StopPercent: stopPercent,
			}
		}
	}
}
