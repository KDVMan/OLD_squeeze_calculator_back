package services_websocket

import (
	"backend/core/models"
	"backend/enums"
	"log"
)

func (websocketService *WebsocketService) broadcastSymbols() {
	symbols, err := websocketService.symbolService.LoadAll()

	if err != nil {
		log.Printf("[broadcastSymbols] failed to load: %v", err)
		return
	}

	broadcastModel := &core_models.BroadcastChannelModel{
		Event: enums.WebsocketEventSymbolCalculatorSymbols,
		Data:  symbols,
	}

	websocketService.BroadcastChan <- broadcastModel
}
