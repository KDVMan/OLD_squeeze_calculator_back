package services_exchange_websocket

import (
	"github.com/adshao/go-binance/v2/futures"
	"log"
	"time"
)

func (exchangeWebsocketService *ExchangeWebsocketService) allMarket() {
	reconnectAttempts := 0
	maxReconnectAttempts := 5

	log.Println("[AllMarket] starting websocket service")

	reconnect := func() {
		if reconnectAttempts < maxReconnectAttempts {
			log.Println("[AllMarket] Attempting to reconnect...")
			time.Sleep(5 * time.Second)
			reconnectAttempts++
			exchangeWebsocketService.allMarket()
		} else {
			log.Println("[AllMarket] Maximum reconnect attempts reached, stopping...")
		}
	}

	handler := func(event futures.WsAllMarketTickerEvent) {
		if err := exchangeWebsocketService.symbolService.UpdateStatistic(event); err != nil {
			log.Println("[AllMarket] failed to update statistic:", err)
		}
	}

	errHandler := func(err error) {
		log.Println("[AllMarket] websocket error:", err)
		reconnect()
	}

	_, _, err := futures.WsAllMarketTickerServe(handler, errHandler)

	if err != nil {
		log.Fatalln("[AllMarket] failed to set up websocket:", err)
	}

	<-exchangeWebsocketService.done
	reconnect()
}
