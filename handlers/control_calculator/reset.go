package handlers_control_calculator

import (
	"backend/core/services/app"
	core_services_logger "backend/core/services/logger"
	"backend/core/services/request"
	core_services_response "backend/core/services/response"
	requests_control_calculator "backend/requests/control_calculator"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func ResetHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := appService.LoggerService.With(
			slog.String("label", "handlers.control_calculator.ResetHandler"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var request requests_control_calculator.ResetRequest

		if err := core_services_request.Decode(w, r, &request, logger); err != nil {
			return
		}

		if err := core_services_request.Validate(w, r, request, logger); err != nil {
			return
		}

		controlCalculatorModel, err := appService.ControlCalculatorService.Reset(request)

		if err != nil {
			message := "failed to reset controlCalculatorModel"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		// logger.Info("controlCalculatorModel loaded", slog.Any("controlCalculatorModel", controlCalculatorModel))
		render.JSON(w, r, controlCalculatorModel)
	}
}
