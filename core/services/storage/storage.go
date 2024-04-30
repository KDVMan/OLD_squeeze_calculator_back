package core_services_storage

import (
	"backend/models/control_symbol"
	"backend/models/exchange"
	"backend/models/init"
	"backend/models/quote"
	"backend/models/result_calculator"
	"backend/models/symbol"
	"backend/models/symbol_calculator"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type StorageService struct {
	DB *gorm.DB
}

func New(path string) (*StorageService, error) {
	const label = "core.services.storage.New"

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	if err = db.Exec("PRAGMA foreign_keys = ON", nil).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	err = db.AutoMigrate(
		&models_exchange.ExchangeLimitModel{},
		&models_init.InitModel{},
		&models_symbol.SymbolModel{},
		&models_symbol_calculator.SymbolCalculatorModel{},
		&models_control_calculator.ControlCalculatorModel{},
		&models_quote.QuoteModel{},
		&models_result_calculator.ResultCalculatorModel{},
		&models_result_calculator.ResultCalculatorDeal{},
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", label, err)
	}

	return &StorageService{DB: db}, nil
}
