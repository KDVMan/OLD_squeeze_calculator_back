package core_services_router

import (
	"backend/core/middlewares/logger"
	"backend/core/services/app"
	"backend/core/services/websocket"
	"backend/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func New(appService *core_services_app.AppService) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(core_middlewares_logger.New(appService.LoggerService))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		services_websocket.HandleConnections(appService.WebsocketService, w, r)
	})

	router.Route("/api", func(r chi.Router) {
		routes.InitRoute(r, appService)
		routes.SymbolRoute(r, appService)
		routes.SymbolCalculatorRoute(r, appService)
		routes.ControlCalculatorRoute(r, appService)
		routes.ResultCalculatorRoute(r, appService)
	})

	return router
}
