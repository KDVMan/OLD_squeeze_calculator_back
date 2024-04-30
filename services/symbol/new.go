package services_symbol

import (
	"backend/core/models"
	"backend/core/services/storage"
)

type SymbolService struct {
	storageService *core_services_storage.StorageService
	broadcastChan  chan *core_models.BroadcastChannelModel
}

func New(storageService *core_services_storage.StorageService) *SymbolService {
	return &SymbolService{
		storageService: storageService,
	}
}

func (symbolService *SymbolService) NewBroadcastChan(broadcastChan chan *core_models.BroadcastChannelModel) {
	symbolService.broadcastChan = broadcastChan
}
