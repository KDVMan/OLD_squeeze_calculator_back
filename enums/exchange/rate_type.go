package enums_exchange

type RateType string

const (
	RateTypeWeight  RateType = "weight"
	RateTypeOrder   RateType = "order"
	RateTypeRequest RateType = "request"
	RateTypeUnknown RateType = "unknown"
)
