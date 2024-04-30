package handlers_result_calculator

import (
	"backend/core/services/app"
	"backend/core/services/logger"
	"backend/core/services/request"
	"backend/core/services/response"
	"backend/requests/result_calculator"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func LoadHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := appService.LoggerService.With(
			slog.String("label", "handlers.result_calculator.LoadHandler"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var request requests_result_calculator.LoadRequest

		if err := core_services_request.Decode(w, r, &request, logger); err != nil {
			return
		}

		if err := core_services_request.Validate(w, r, request, logger); err != nil {
			return
		}

		resultCalculatorModel, err := appService.ResultCalculatorService.Load(request)

		if err != nil {
			message := "failed to load resultCalculatorModel"
			logger.Error(message, core_services_logger.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, core_services_response.Error(message))

			return
		}

		// log.Println("resultCalculatorModel loaded", slog.Any("resultCalculatorModel", core_services_helper.ModelToJson(resultCalculatorModel)))
		render.JSON(w, r, resultCalculatorModel)
	}
}
