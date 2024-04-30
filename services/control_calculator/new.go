package services_control_calculator

import (
	"backend/core/services/storage"
)

type ControlCalculatorService struct {
	storageService *core_services_storage.StorageService
}

func New(storageService *core_services_storage.StorageService) *ControlCalculatorService {
	return &ControlCalculatorService{
		storageService: storageService,
	}
}
