package routes

import (
	"backend/core/services/app"
	"backend/handlers/init"
	"github.com/go-chi/chi/v5"
)

func InitRoute(router chi.Router, appService *core_services_app.AppService) {
	router.Get("/init/load", handlers_init.LoadHandler(appService))
	router.Post("/init/update", handlers_init.UpdateHandler(appService))
}
