package handlers_control_calculator

import (
	"backend/core/services/app"
	"backend/variables/calculator"
	"github.com/go-chi/render"
	"net/http"
)

func StopHandler(appService *core_services_app.AppService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		variables_calculator.Stop = true
		render.JSON(w, r, nil)
	}
}
