package services_init

import (
	"backend/models/init"
	"backend/requests/init"
	"fmt"
)

func (initService *InitService) Update(request requests_init.UpdateRequest) (*models_init.InitModel, error) {
	const label = "services.init.Update"

	initModel, err := initService.Load()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	initModel.Symbol = request.Symbol
	initModel.Instrument = request.Instrument

	result := initService.storageService.DB.Save(initModel)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", label, result.Error)
	}

	return initModel, nil
}
