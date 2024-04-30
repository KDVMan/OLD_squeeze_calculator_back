package handlers_init

import (
	"backend/core/services/app"
	"backend/core/services/logger"
	"backend/core/services/request"
	"backend/core/services/response"
	"backend/requests/init"
	variables_calculator "backend/variables/calculator"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func UpdateHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := appService.LoggerService.With(
			slog.String("label", "handlers.init.UpdateHandler"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		variables_calculator.Stop = true

		var request requests_init.UpdateRequest

		if err := core_services_request.Decode(w, r, &request, logger); err != nil {
			return
		}

		if err := core_services_request.Validate(w, r, request, logger); err != nil {
			return
		}

		initModel, err := appService.InitService.Update(request)

		if err != nil {
			message := "failed to update initModel"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		logger.Info("initModel updated", slog.Any("initModel", initModel))
		render.JSON(w, r, initModel)
	}
}
