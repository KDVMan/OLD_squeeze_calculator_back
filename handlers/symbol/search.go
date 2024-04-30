package handlers_symbol

import (
	"backend/core/services/app"
	"backend/core/services/logger"
	"backend/core/services/request"
	"backend/core/services/response"
	"backend/requests/symbol"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func SearchHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := appService.LoggerService.With(
			slog.String("label", "handlers.symbol.SearchHandler"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var request requests_symbol.SearchRequest

		if err := core_services_request.Decode(w, r, &request, logger); err != nil {
			return
		}

		if err := core_services_request.Validate(w, r, request, logger); err != nil {
			return
		}

		symbolModel, err := appService.SymbolService.Search(request)

		if err != nil {
			message := "failed to load symbolModel"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		// logger.Info("symbolModel loaded", slog.Any("symbolModel", symbolModel))
		render.JSON(w, r, symbolModel)
	}
}
