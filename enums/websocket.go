package enums

type WebsocketEvent string

const (
	WebsocketEventSymbolCalculatorSymbols WebsocketEvent = "symbolCalculatorSymbols"
	WebsocketEventExchangeLimits          WebsocketEvent = "exchangeLimits"
	WebsocketEventCalculator              WebsocketEvent = "calculator"
)
