package routes

import (
	"backend/core/services/app"
	"backend/handlers/symbol"
	"github.com/go-chi/chi/v5"
)

func SymbolRoute(router chi.Router, appService *core_services_app.AppService) {
	router.Post("/symbol/search", handlers_symbol.SearchHandler(appService))
	router.Get("/symbol/download", handlers_symbol.DownloadHandler(appService))
}
