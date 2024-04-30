package enums_result_calculator

import "github.com/go-playground/validator/v10"

type SortColumn string

const (
	SortColumnPercentIn            SortColumn = "percentIn"
	SortColumnPercentOut           SortColumn = "percentOut"
	SortColumnBind                 SortColumn = "bind"
	SortColumnStopTime             SortColumn = "stopTime"
	SortColumnStopPercent          SortColumn = "stopPercent"
	SortColumnProfitPercent        SortColumn = "profitPercent"
	SortColumnAverageProfitPercent SortColumn = "averageProfitPercent"
	SortColumnTotal                SortColumn = "total"
	SortColumnTotalStops           SortColumn = "totalStops"
	SortColumnTotalTakes           SortColumn = "totalTakes"
	SortColumnCoefficient          SortColumn = "coefficient"
	SortColumnRatio                SortColumn = "ratio"
	SortColumnWinRate              SortColumn = "winRate"
	SortColumnMaxTimeDeal          SortColumn = "maxTimeDeal"
	SortColumnAverageTimeDeal      SortColumn = "averageTimeDeal"
	SortColumnMaxDrawdown          SortColumn = "maxDrawdown"
	SortColumnAverageDrawdown      SortColumn = "averageDrawdown"
	SortColumnDrawdownProfitRatio  SortColumn = "drawdownProfitRatio"
	SortColumnScore                SortColumn = "score"
)

func SortColumnValidate(field validator.FieldLevel) bool {
	if enum, ok := field.Field().Interface().(SortColumn); ok {
		return enum.SortColumnValid()
	}

	return false
}

func (enum SortColumn) SortColumnValid() bool {
	switch enum {
	case SortColumnPercentIn, SortColumnPercentOut, SortColumnBind, SortColumnStopTime, SortColumnStopPercent, SortColumnProfitPercent, SortColumnAverageProfitPercent,
		SortColumnTotal, SortColumnTotalStops, SortColumnTotalTakes, SortColumnCoefficient, SortColumnRatio, SortColumnWinRate, SortColumnMaxTimeDeal, SortColumnAverageTimeDeal,
		SortColumnMaxDrawdown, SortColumnAverageDrawdown, SortColumnDrawdownProfitRatio, SortColumnScore:
		return true
	default:
		return false
	}
}
