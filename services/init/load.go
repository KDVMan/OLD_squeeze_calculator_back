package services_init

import (
	"backend/models/init"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (initService *InitService) Load() (*models_init.InitModel, error) {
	const label = "services.init.Load"
	var initModel *models_init.InitModel

	if err := initService.storageService.DB.First(&initModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			initModel = models_init.LoadDefault()

			if err := initService.storageService.DB.Create(initModel).Error; err != nil {
				return nil, fmt.Errorf("%s: %w", label, err)
			}

			return initModel, nil
		}

		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return initModel, nil
}
