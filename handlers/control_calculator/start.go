package handlers_control_calculator

import (
	"backend/core/services/app"
	"backend/core/services/logger"
	"backend/core/services/request"
	"backend/core/services/response"
	"backend/enums/symbol"
	"backend/requests/control_calculator"
	"backend/services/calculator"
	"backend/variables/calculator"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func StartHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := appService.LoggerService.With(
			slog.String("label", "handlers.control_calculator.StartHandler"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var request requests_control_calculator.StartRequest

		if err := core_services_request.Decode(w, r, &request, logger); err != nil {
			return
		}

		if err := core_services_request.Validate(w, r, request, logger); err != nil {
			return
		}

		symbolModel, err := appService.SymbolService.Load(request.Symbol, enums_symbol.SymbolStatusActive)

		if err != nil {
			message := "Failed to load symbol"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		controlCalculatorModel, err := appService.ControlCalculatorService.Update(request)

		if err != nil {
			message := "[2] Failed to start calculator"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		if err = appService.ResultCalculatorService.Clear(controlCalculatorModel.Symbol, controlCalculatorModel.TradeDirection, controlCalculatorModel.Interval); err != nil {
			message := "Failed to clear result calculator"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		variables_calculator.Stop = false
		go services_calculator.Start(appService, logger, controlCalculatorModel, symbolModel, &request)

		render.JSON(w, r, nil)
	}
}
