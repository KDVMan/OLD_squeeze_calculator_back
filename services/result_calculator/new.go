package services_result_calculator

import (
	"backend/core/services/storage"
)

type ResultCalculatorService struct {
	storageService *core_services_storage.StorageService
}

func New(storageService *core_services_storage.StorageService) *ResultCalculatorService {
	return &ResultCalculatorService{
		storageService: storageService,
	}
}
