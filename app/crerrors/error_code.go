package crerrors

const (
	// Bad Request
	BadRequestErrorCode     = "BAD_REQUEST_ERROR"
	InternalServerErrorCode = "INTERNAL_SERVER_ERROR"
	CodeNotFound            = "NOT_FOUND"

	// CodeBadRequest : ...
	CodeBadRequest = "BAD_REQUEST"

	// CodeServerError : ...
	CodeServerError = "SERVER_ERROR"

	// CodeNoDataFound : represent there was no data available in db
	CodeNoDataFound = "NO_DATA_FUND"

	// CodeUnauthorized : access denied for the used
	CodeUnauthorized = "UNAUTHORIZED"

	// CodeEmptyPayload : payload to enqueue was empty
	CodeEmptyPayload = "EMPTY_PAYLOAD"

	// CodeRuntimeError : ...
	CodeRuntimeError = "RUNTIME_ERROR"

	// CodeDatabaseError : database operation failed
	CodeDatabaseError = "DATABASE_ERROR"

	CodeRecordsAlreadyExist = "RECORD_ALREADY_EXIST"

	CodeRecordNotFound = "RECORD_NOT_FOUND"

	// CodeBindingFailed : data binding with struct has failed
	CodeBindingFailed = "FAILED_TO_BIND_DATA_TO_STRUCT"

	// CodeIllegalArgument : argument is of invalid type
	CodeIllegalArgument = "ILLEGAL_ARGUMENT"

	// CodeValidationError : validation error for a struct
	CodeValidationError = "VALIDATION_ERROR"

	CodeRecordDuplicateEntry = "RECORD_DUPLICATE_ENTRY"
)
