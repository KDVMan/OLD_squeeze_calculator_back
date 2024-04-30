package enums_exchange

type RateInterval string

const (
	RateIntervalSecond  RateInterval = "second"
	RateIntervalMinute  RateInterval = "minute"
	RateIntervalHour    RateInterval = "hour"
	RateIntervalDay     RateInterval = "day"
	RateIntervalUnknown RateInterval = "unknown"
)
