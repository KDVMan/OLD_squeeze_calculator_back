package core_services_helper

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func ConvertFloat(value string, valueDefault float64, base int) float64 {
	if result, err := strconv.ParseFloat(value, base); err == nil {
		return result
	}

	return valueDefault
}

func GetPercentFromMinMax(min float64, max float64, fix int) float64 {
	if min == 0 {
		return 0
	}

	result := ((max / min) * 100) - 100

	if fix > 0 {
		return Round(result, fix)
	}

	return result
}

func GetRandomInt(min int64, max int64, step int64) int64 {
	if step <= 0 {
		panic("step must be positive")
	}

	if min > max {
		panic("min cannot be greater than max")
	}

	numSteps := (max-min)/step + 1
	randomStep := rand.Int63n(numSteps)

	return min + randomStep*step
}

func GetRandomFloatByInt(min float64, max float64, step float64) float64 {
	accuracy := CalculateAccuracy(step)
	minInt := int64(min * float64(accuracy))
	maxInt := int64(max * float64(accuracy))
	stepInt := int64(step * float64(accuracy))
	rangeInt := maxInt - minInt
	steps := rangeInt / stepInt
	randomStep := rand.Int63n(steps + 1)

	return float64(minInt+randomStep*stepInt) / float64(accuracy)
}

func GetRangeFloatByInt(min float64, max float64, step float64) (int64, int64, int64, int64) {
	accuracy := CalculateAccuracy(step)
	minInt := int64(min * float64(accuracy))
	maxInt := int64(max * float64(accuracy))
	stepInt := int64(step * float64(accuracy))

	return minInt, maxInt, stepInt, accuracy
}

func CalculateAccuracy(step float64) int64 {
	stepStr := fmt.Sprintf("%g", step)
	decimalPlaces := 0

	if strings.Contains(stepStr, ".") {
		decimalPlaces = len(stepStr) - strings.Index(stepStr, ".") - 1
	}

	accuracy := int64(math.Pow10(decimalPlaces))

	if accuracy > 10 {
		return accuracy
	} else {
		return 10
	}
}

func CalculateStandardDeviation(values []float64, mean float64) float64 {
	var sum float64

	if len(values) == 0 {
		return 0
	}

	for _, v := range values {
		sum += math.Pow(v-mean, 2)
	}

	return math.Sqrt(sum / float64(len(values)))
}

func CalculateAverage(values []float64) float64 {
	sum := 0.0

	if len(values) == 0 {
		return 0
	}

	for _, v := range values {
		sum += v
	}

	return sum / float64(len(values))
}

func Round(value float64, decimal int) float64 {
	if decimal == 0 {
		return math.Round(value)
	}

	multiplier := math.Pow(10, float64(decimal))

	return math.Round(value*multiplier) / multiplier
}

func Floor(value float64, decimal int) float64 {
	if decimal == 0 {
		return math.Floor(value)
	}

	multiplier := math.Pow(10, float64(decimal))

	return math.Floor(value*multiplier) / multiplier
}
