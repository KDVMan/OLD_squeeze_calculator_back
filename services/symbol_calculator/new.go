package services_symbol_calculator

import (
	"backend/core/services/storage"
)

type SymbolCalculatorService struct {
	storageService *core_services_storage.StorageService
}

func New(storageService *core_services_storage.StorageService) *SymbolCalculatorService {
	return &SymbolCalculatorService{
		storageService: storageService,
	}
}
