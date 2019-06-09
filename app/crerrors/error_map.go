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

	CodeRecordNotFound :{
		publicCode:    CodeNotFound,
		statusCode:    http.StatusBadRequest,
	},
	CodeDatabaseError:{
		publicCode:    CodeDatabaseError,
		statusCode:    http.StatusInternalServerError,
	},
	SecretExpired:{
		publicCode:    SecretExpired,
		statusCode:    http.StatusBadRequest,
	},
	SecretLimitReached:{
		publicCode:    SecretLimitReached,
		statusCode:    http.StatusBadRequest,
	},
}

func getPublicData(code string) IFields {
	data, _ := errorMap[code]
	return &data
}
