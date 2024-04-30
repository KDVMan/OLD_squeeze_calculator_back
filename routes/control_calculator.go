package routes

import (
	"backend/core/services/app"
	"backend/handlers/control_calculator"
	"github.com/go-chi/chi/v5"
)

func ControlCalculatorRoute(router chi.Router, appService *core_services_app.AppService) {
	router.Post("/control_calculator/load", handlers_control_calculator.LoadHandler(appService))
	router.Post("/control_calculator/start", handlers_control_calculator.StartHandler(appService))
	router.Get("/control_calculator/stop", handlers_control_calculator.StopHandler(appService))
	router.Post("/control_calculator/reset", handlers_control_calculator.ResetHandler(appService))
}
