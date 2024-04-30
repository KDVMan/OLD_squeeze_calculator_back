package core_services_request

import (
	"backend/core/services/logger"
	"backend/core/services/response"
	"backend/enums"
	"backend/enums/result_calculator"
	"backend/enums/symbol_calculator"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	if err := validate.RegisterValidation("instrument", enums.InstrumentValidate); err != nil {
		return
	}

	if err := validate.RegisterValidation("symbolCalculatorSortColumn", enums_symbol_calculator.SortColumnValidate); err != nil {
		return
	}

	if err := validate.RegisterValidation("resultCalculatorSortColumn", enums_result_calculator.SortColumnValidate); err != nil {
		return
	}

	if err := validate.RegisterValidation("sortDirection", enums.SortDirectionValidate); err != nil {
		return
	}

	if err := validate.RegisterValidation("tradeDirection", enums.TradeDirectionValidate); err != nil {
		return
	}

	if err := validate.RegisterValidation("interval", enums.IntervalValidate); err != nil {
		return
	}

	if err := validate.RegisterValidation("bind", enums.BindValidate); err != nil {
		return
	}

	if err := validate.RegisterValidation("algorithm", enums.AlgorithmValidate); err != nil {
		return
	}
}

func Validate(w http.ResponseWriter, r *http.Request, request interface{}, logger *slog.Logger) error {
	if err := validate.Struct(request); err != nil {
		var validateError validator.ValidationErrors

		if errors.As(err, &validateError) {
			logger.Error("validation failed", core_services_logger.Err(err))

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, getErrors(validateError))

			return err
		}

		message := "unknown validation error"
		logger.Error(message, core_services_logger.Err(err))

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, core_services_response.Error(message))

		return err
	}

	return nil
}

func getErrors(errors validator.ValidationErrors) core_services_response.Response {
	var messages []string

	for _, err := range errors {
		switch err.ActualTag() {
		case "required":
			messages = append(messages, fmt.Sprintf("field %s is a required field", err.Field()))
		case "alphanum":
			messages = append(messages, fmt.Sprintf("field %s is not valid, must be alphanumeric", err.Field()))
		case "uppercase":
			messages = append(messages, fmt.Sprintf("field %s is not valid, must be uppercase", err.Field()))
		default:
			messages = append(messages, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return core_services_response.Response{
		Status: core_services_response.StatusError,
		Error:  strings.Join(messages, ", "),
	}
}
