package crerrors

const (
	// InternalErrorDetails : index tp hold the internal error details
	InternalErrorDetails = "internal_error_details"

	// PublicErrorDetails : index tp hold the public error details
	PublicErrorDetails = "public_error_details"

	// Code : index to hold the error code
	Code = "code"

	// SubCode will refer to the sub error code through which the error can be distinguished
	SubCode = "sub_code"

	// Message : index to hold the error message
	Message = "message"

	// Errors : index to hold the entity specific errors
	Errors = "errors"

	DebugMessage = "debug_message"
)
