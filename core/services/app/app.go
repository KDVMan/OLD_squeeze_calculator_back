package core_services_app

import (
	"backend/core/services/config"
	"backend/core/services/logger"
	"backend/core/services/storage"
	"backend/core/services/websocket"
	"backend/services/control_calculator"
	"backend/services/exchange"
	"backend/services/exchange_limit"
	"backend/services/exchange_websocket"
	"backend/services/init"
	"backend/services/quote"
	"backend/services/result_calculator"
	"backend/services/symbol"
	"backend/services/symbol_calculator"
	"log/slog"
	"os"
)

type AppService struct {
	ConfigService            *core_services_config.ConfigService
	LoggerService            *slog.Logger
	StorageService           *core_services_storage.StorageService
	InitService              *services_init.InitService
	SymbolService            *services_symbol.SymbolService
	SymbolCalculatorService  *services_symbol_calculator.SymbolCalculatorService
	ControlCalculatorService *services_control_calculator.ControlCalculatorService
	ExchangeLimitService     *services_exchange_limit.ExchangeLimitService
	ExchangeService          *services_exchange.ExchangeService
	ExchangeWebsocketService *services_exchange_websocket.ExchangeWebsocketService
	QuoteService             *services_quote.QuoteService
	ResultCalculatorService  *services_result_calculator.ResultCalculatorService
	WebsocketService         *services_websocket.WebsocketService
}

func New() *AppService {
	configService := core_services_config.New()
	loggerService := core_services_logger.New(configService.Env)
	storageService := initStorage(loggerService, configService.StoragePath)
	initService := services_init.New(storageService)
	symbolService := services_symbol.New(storageService)
	symbolCalculatorService := services_symbol_calculator.New(storageService)
	controlCalculatorService := services_control_calculator.New(storageService)
	exchangeLimitService := services_exchange_limit.New(storageService)
	exchangeService := services_exchange.New(exchangeLimitService)
	exchangeWebsocketService := services_exchange_websocket.New(symbolService)
	quoteService := services_quote.New(storageService, exchangeService)
	resultCalculatorService := services_result_calculator.New(storageService)
	websocketService := services_websocket.New(symbolService, exchangeLimitService)

	symbolService.NewBroadcastChan(websocketService.BroadcastChan)
	exchangeLimitService.NewBroadcastChan(websocketService.BroadcastChan)

	return &AppService{
		ConfigService:            configService,
		LoggerService:            loggerService,
		StorageService:           storageService,
		InitService:              initService,
		SymbolService:            symbolService,
		SymbolCalculatorService:  symbolCalculatorService,
		ControlCalculatorService: controlCalculatorService,
		ExchangeLimitService:     exchangeLimitService,
		ExchangeService:          exchangeService,
		ExchangeWebsocketService: exchangeWebsocketService,
		QuoteService:             quoteService,
		ResultCalculatorService:  resultCalculatorService,
		WebsocketService:         websocketService,
	}
}

func initStorage(loggerService *slog.Logger, path string) *core_services_storage.StorageService {
	storageService, err := core_services_storage.New(path)

	if err != nil {
		loggerService.Error("failed to init storage", core_services_logger.Err(err))
		os.Exit(1)
	}

	return storageService
}
