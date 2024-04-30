package services_quote

import (
	"backend/core/services/storage"
	services_exchange "backend/services/exchange"
)

type QuoteService struct {
	storageService  *core_services_storage.StorageService
	exchangeService *services_exchange.ExchangeService
}

func New(storageService *core_services_storage.StorageService, exchangeService *services_exchange.ExchangeService) *QuoteService {
	return &QuoteService{
		storageService:  storageService,
		exchangeService: exchangeService,
	}
}
