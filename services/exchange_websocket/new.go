package services_exchange_websocket

import (
	"backend/services/symbol"
	"log"
)

type ExchangeWebsocketService struct {
	symbolService *services_symbol.SymbolService
	done          chan struct{}
}

func New(symbolService *services_symbol.SymbolService) *ExchangeWebsocketService {
	return &ExchangeWebsocketService{
		symbolService: symbolService,
		done:          make(chan struct{}),
	}
}

func (exchangeWebsocketService *ExchangeWebsocketService) Start() {
	// go exchangeWebsocketService.allMarket()
}

func (exchangeWebsocketService *ExchangeWebsocketService) Stop() {
	log.Println("Stopping exchange websocket service")
	close(exchangeWebsocketService.done)
}
