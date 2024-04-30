package services_init

import (
	"backend/core/services/storage"
)

type InitService struct {
	storageService *core_services_storage.StorageService
}

func New(storageService *core_services_storage.StorageService) *InitService {
	return &InitService{
		storageService: storageService,
	}
}
