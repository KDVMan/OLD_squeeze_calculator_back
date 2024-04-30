package routes

import (
	"backend/core/services/app"
	"backend/handlers/symbol_calculator"
	"github.com/go-chi/chi/v5"
)

func SymbolCalculatorRoute(router chi.Router, appService *core_services_app.AppService) {
	router.Get("/symbol_calculator/load", handlers_symbol_calculator.LoadHandler(appService))
	router.Post("/symbol_calculator/update", handlers_symbol_calculator.UpdateHandler(appService))
}
