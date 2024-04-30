package enums

import "github.com/go-playground/validator/v10"

type Instrument string

const (
	InstrumentCalculator Instrument = "calculator"
	InstrumentMonitor    Instrument = "monitor"
)

func InstrumentValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Instrument); ok {
		return enum.InstrumentValid()
	}

	return false
}

func (enum Instrument) InstrumentValid() bool {
	switch enum {
	case InstrumentCalculator, InstrumentMonitor:
		return true
	default:
		return false
	}
}
