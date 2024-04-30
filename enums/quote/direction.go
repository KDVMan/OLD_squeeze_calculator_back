package enums_quote

import "github.com/go-playground/validator/v10"

type Direction string

const (
	DirectionNone Direction = "none"
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

func DirectionValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(Direction); ok {
		return enum.DirectionValid()
	}

	return false
}

func (enum Direction) DirectionValid() bool {
	switch enum {
	case DirectionNone, DirectionUp, DirectionDown:
		return true
	default:
		return false
	}
}
