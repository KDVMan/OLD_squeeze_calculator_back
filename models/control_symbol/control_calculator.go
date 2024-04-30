package models_control_calculator

import (
	"backend/enums"
	"backend/models"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ControlCalculatorModel struct {
	models.BaseModel
	Symbol          string               `gorm:"uniqueIndex:unique_control_calculator;not null" json:"symbol"`
	TradeDirection  enums.TradeDirection `gorm:"uniqueIndex:unique_control_calculator;not null" json:"tradeDirection"`
	Interval        enums.Interval       `gorm:"uniqueIndex:unique_control_calculator;not null" json:"interval"`
	TimeFrom        int64                `json:"timeFrom"`
	TimeTo          int64                `json:"timeTo"`
	OncePerCandle   bool                 `json:"oncePerCandle"`
	Bind            []enums.Bind         `gorm:"-" json:"bind"`
	BindJson        string               `gorm:"type:text" json:"-"`
	PercentInFrom   float64              `json:"percentInFrom"`
	PercentInTo     float64              `json:"percentInTo"`
	PercentInStep   float64              `json:"percentInStep"`
	PercentOutFrom  float64              `json:"percentOutFrom"`
	PercentOutTo    float64              `json:"percentOutTo"`
	PercentOutStep  float64              `json:"percentOutStep"`
	StopTime        bool                 `json:"stopTime"`
	StopTimeFrom    int64                `json:"stopTimeFrom"`
	StopTimeTo      int64                `json:"stopTimeTo"`
	StopTimeStep    int64                `json:"stopTimeStep"`
	StopPercent     bool                 `json:"stopPercent"`
	StopPercentFrom float64              `json:"stopPercentFrom"`
	StopPercentTo   float64              `json:"stopPercentTo"`
	StopPercentStep float64              `json:"stopPercentStep"`
	Algorithm       enums.Algorithm      `json:"algorithm"`
	Iterations      int                  `json:"iterations"`
}

func (ControlCalculatorModel) TableName() string {
	return "controls_calculators"
}

func (model *ControlCalculatorModel) BeforeSave(tx *gorm.DB) (err error) {
	bindData, err := json.Marshal(model.Bind)

	if err != nil {
		return err
	}

	model.BindJson = string(bindData)

	return nil
}

func (model *ControlCalculatorModel) AfterFind(tx *gorm.DB) (err error) {
	if model.BindJson != "" {
		err := json.Unmarshal([]byte(model.BindJson), &model.Bind)

		if err != nil {
			return err
		}
	}

	return nil
}

func LoadDefault(symbol string, tradeDirection enums.TradeDirection, interval enums.Interval) *ControlCalculatorModel {
	currentTime := time.Now()
	start := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).AddDate(0, 0, -10)
	end := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())

	model := &ControlCalculatorModel{
		Symbol:          symbol,
		TradeDirection:  tradeDirection,
		Interval:        interval,
		TimeFrom:        start.UnixMilli(),
		TimeTo:          end.UnixMilli(),
		OncePerCandle:   false,
		Bind:            enums.BindValues(),
		PercentInFrom:   1,
		PercentInTo:     5,
		PercentInStep:   0.1,
		PercentOutFrom:  0.1,
		PercentOutTo:    3,
		PercentOutStep:  0.1,
		StopTime:        true,
		StopTimeFrom:    1,
		StopTimeTo:      92,
		StopTimeStep:    1,
		StopPercent:     true,
		StopPercentFrom: 0.1,
		StopPercentTo:   10,
		StopPercentStep: 0.1,
		Algorithm:       enums.AlgorithmRandom,
		Iterations:      50000,
	}

	switch interval {
	case enums.Interval3m:
		model.PercentInTo = 7
		model.PercentOutTo = 4
	case enums.Interval5m:
		model.PercentInTo = 9
		model.PercentOutTo = 6
	case enums.Interval15m:
		model.PercentInTo = 12
		model.PercentOutTo = 8
	case enums.Interval1h:
		model.PercentInTo = 20
		model.PercentOutTo = 12
	default:
	}

	return model
}
