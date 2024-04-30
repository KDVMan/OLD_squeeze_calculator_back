package routes

import (
	"backend/core/services/app"
	handlers_result_calculator "backend/handlers/result_calculator"
	"github.com/go-chi/chi/v5"
)

func ResultCalculatorRoute(router chi.Router, appService *core_services_app.AppService) {
	router.Post("/result_calculator/load", handlers_result_calculator.LoadHandler(appService))
}
