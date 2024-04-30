package core_services_request

import (
	"backend/core/services/logger"
	"backend/core/services/response"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func Decode(w http.ResponseWriter, r *http.Request, request interface{}, logger *slog.Logger) error {
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		message := "failed to decode request body"
		logger.Error(message, core_services_logger.Err(err))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, core_services_response.Error(message))

		return err
	}

	return nil
}
