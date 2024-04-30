package services_websocket

import (
	"backend/core/models"
	"backend/enums"
	"log"
)

func (websocketService *WebsocketService) broadcastExchangeLimits() {
	limits, err := websocketService.exchangeLimitService.Load()

	if err != nil {
		log.Printf("[broadcastExchangeLimits] failed to load: %v", err)
		return
	}

	broadcastModel := &core_models.BroadcastChannelModel{
		Event: enums.WebsocketEventExchangeLimits,
		Data:  limits,
	}

	websocketService.BroadcastChan <- broadcastModel
}
