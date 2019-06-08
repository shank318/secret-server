package crerrors

import "net/http"

var errorMap = map[string]fields{
	BadRequestErrorCode: {
		publicCode:    BadRequestErrorCode,
		publicMessage: BadRequestErrorMessage,
		statusCode:    http.StatusBadRequest,
	},

	InternalServerErrorCode: {
		publicCode:    InternalServerErrorCode,
		publicMessage: InternalServerErrorMessage,
		statusCode:    http.StatusInternalServerError,
	},
}

func getPublicData(code string) IFields {
	data, _ := errorMap[code]
	return &data
}
