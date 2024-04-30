package services_exchange_limit

import (
	"backend/core/models"
	"backend/core/services/storage"
)

type ExchangeLimitService struct {
	storageService *core_services_storage.StorageService
	broadcastChan  chan *core_models.BroadcastChannelModel
}

func New(storageService *core_services_storage.StorageService) *ExchangeLimitService {
	return &ExchangeLimitService{
		storageService: storageService,
	}
}

func (exchangeLimitService *ExchangeLimitService) NewBroadcastChan(broadcastChan chan *core_models.BroadcastChannelModel) {
	exchangeLimitService.broadcastChan = broadcastChan
}
