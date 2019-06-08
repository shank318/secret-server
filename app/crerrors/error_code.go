package crerrors

const (
	// Bad Request
	BadRequestErrorCode     = "BAD_REQUEST_ERROR"
	InternalServerErrorCode = "INTERNAL_SERVER_ERROR"
	CodeNotFound            = "NOT_FOUND"

	// CodeBadRequest : ...
	CodeBadRequest = "BAD_REQUEST"

	// CodeQueueError : queue operation failed
	CodeQueueError = "QUEUE_ERROR"

	// CodeServerError : ...
	CodeServerError = "SERVER_ERROR"

	// CodeNoDataFound : represent there was no data available in db
	CodeNoDataFound = "NO_DATA_FUND"

	// CodeIllegalState : trying to operate on invalid state
	CodeIllegalState = "ILLEGAL_STATE"

	// CodeUnauthorized : access denied for the used
	CodeUnauthorized = "UNAUTHORIZED"

	// CodeEmptyPayload : payload to enqueue was empty
	CodeEmptyPayload = "EMPTY_PAYLOAD"

	// CodeRuntimeError : ...
	CodeRuntimeError = "RUNTIME_ERROR"

	// CodeCriticalError : something very wrong is happening
	CodeCriticalError = "CRITICAL_ERROR"

	// CodeUnknownQueue : provided queue name is unknown
	CodeUnknownQueue = "UNKNOWN_QUEUE"

	// CodeDatabaseError : database operation failed
	CodeDatabaseError = "DATABASE_ERROR"

	CodeRecordsAlreadyExist = "RECORD_ALREADY_EXIST"

	CodeRecordNotFound = "RECORD_NOT_FOUND"

	// CodeBindingFailed : data binding with struct has failed
	CodeBindingFailed = "FAILED_TO_BIND_DATA_TO_STRUCT"

	// Governor error
	GovernorError = "GOVERNOR_ERROR"

	// CodeIllegalArgument : argument is of invalid type
	CodeIllegalArgument = "ILLEGAL_ARGUMENT"

	// CodeValidationError : validation error for a struct
	CodeValidationError = "VALIDATION_ERROR"

	// CodeRequestModeInvalid : request mode `X-mode` was invalid
	CodeRequestModeInvalid = "REQUEST_MODE_INVALID"

	// CodeCredentialsNotFound : credentials were not found for request
	CodeCredentialsNotFound = "CREDENTIALS_NOT_FOUND"

	// CodeFailedToFetchMessage : failed to fetch message from queue
	CodeFailedToFetchMessage = "FAILED_TO_FETCH_MESSAGE_FROM_QUEUE"

	// CodeServiceRequestFailed : request to the request service failed
	CodeServiceRequestFailed = "SERVICE_REQUEST_FAILED"

	// CodeServiceRequestFailed : request to the request service failed
	CodeRequestURLIncorrect = "REQUEST_URL_INCORRECT"

	// CodeStateTransitionFailed : trying to operate on invalid state
	CodeStateTransitionFailed = "STATE_TRANSITION_FAILED"

	// CodeFailedToEnqueueMessage : failed to add message to the queue
	CodeFailedToEnqueueMessage = "FAILED_TO_ENQUEUE_MESSAGE"

	// CodeFormatConversionFailed : data format conversion failed
	CodeFormatConversionFailed = "FORMAT_CONVERSION_FAILED"

	// CodeFailedToEstablishConnection : failed to establish connection with providers
	CodeFailedToEstablishConnection = "FAILED_TO_ESTABLISH_CONNECTION"

	// *** Mutex related error codes *** //
	CodeAcquiringMutexFailed = "ACQUIRING_MUTEX_FAILED"

	CodeMutexAlreadyReleased = "MUTEX_ALREADY_RELEASED"

	CodeRzpIDLengthMismatchError = "RZPID_LENGTH_MISMATCH"
)
