package models_calculator

import "backend/enums"

type CalculatorOptimizationModel struct {
	Bind        enums.Bind `json:"bind"`
	PercentIn   float64    `json:"percentIn"`
	PercentOut  float64    `json:"percentOut"`
	StopTime    int64      `json:"stopTime"`
	StopPercent float64    `json:"stopPercent"`
}
