package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    DataNil `json:"data"`
}

type DataNil struct {
}

func FormatErrorValidation(err error) ErrorResponse {
	var messageBody string
	var errorResponse ErrorResponse
	for _, v := range err.(validator.ValidationErrors) {
		switch v.Tag() {
		case "required":
			messageBody = strings.ToLower(v.Field()) + ` cannot be null`
		default:
			messageBody = v.Error()
		}

		errorResponse.Status = "Bad Request"
		errorResponse.Message = messageBody
		// errorResponse.Data = interfaces{}
		return errorResponse

	}
	return errorResponse
}
