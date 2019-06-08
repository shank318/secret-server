package crerrors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"secret-server/app/metric"
	"secret-server/app/logger"
)

type IError interface {
	Log() IError
	GetInternalError() error
	Error() string
	Message() string
	SubCode() string
	ErrorCode() string
	StatusCode() int
	PublicMessage() string
	PublicErrorCode() string
	WithField(string, interface{}) IError
	WithFields(map[string]interface{}) IError
	WithInternalMessage(string) IError
	WithSubCode(interface{}) IError
	IsSubCode(string) bool
	InternalErrorDetails() map[string]interface{}
	PublicErrorDetails() map[string]interface{}
	ErrorData() []map[string]interface{}
	SetErrorData([]map[string]interface{}) IError
	GetTraceData() map[string]interface{}
}

type cRError struct {
	IError
	err                error
	errorData          []map[string]interface{}
	errorCode          string
	subCode            string
	logData            map[string]interface{}
	errorType          string
	publicErrorDetails IFields
	ctx                *gin.Context
}

// NewCrError will create a new error given error code
// We cannot have zero value (nil) for logData since
// when we are fetching the logData, we also need
// the ability to directly add more keys to it.
// But if it's nil, we would have to initialize
// it again separately while fetching the logData.
func NewCrError(ctx *gin.Context, code string, err error) IError {
	rzpErr := &cRError{
		err:                err,
		errorCode:          code,
		publicErrorDetails: getPublicData(code),
		logData:            map[string]interface{}{},
		ctx:                ctx,
	}
	metric.ErrorCounter.WithLabelValues(rzpErr.ErrorCode()).Inc()
	return rzpErr
}

func CopyRzpError(ctx *gin.Context, code string, err IError) IError {
	rzpErr := &cRError{
		err:                err.GetInternalError(),
		errorCode:          code,
		publicErrorDetails: getPublicData(code),
		logData:            map[string]interface{}{},
		ctx:                ctx,
	}
	return rzpErr
}

// Message will give the internal error message
// Message can be nil in case of custom error where we don't have any error
// which is generated while running the application
func (e *cRError) Message() string {
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}

// ErrorCode gives the error code back
func (e *cRError) ErrorCode() string {
	return e.errorCode
}

// SubCode will give the sub code of the error
func (e *cRError) SubCode() string {
	return e.subCode
}

// GetInternalError will give the error interface
func (e *cRError) GetInternalError() error {
	return e.err
}

// Error will give the error string
func (e *cRError) Error() string {
	return e.err.Error()
}

// StatusCode will get the status code
func (e *cRError) StatusCode() int {
	return e.publicErrorDetails.StatusCode()
}

// PublicMessage will give the public message is set else will return the error message
func (e *cRError) PublicMessage() string {
	return e.publicErrorDetails.Message()
}

// PublicMessage will give the public message is set else will return the error message
func (e *cRError) PublicErrorCode() string {
	return e.publicErrorDetails.Code()
}

func (e *cRError) Level() string {
	return e.publicErrorDetails.Level()
}

// WithFields sets the error logData
func (e *cRError) WithFields(data map[string]interface{}) IError {
	for key, value := range data {
		e.logData[key] = value
	}

	return e
}

func (e *cRError) SetErrorData(errorData []map[string]interface{}) IError {
	e.errorData = append(e.errorData, errorData...)

	return e
}

// WithSubCode will set the sub code of the error through which the error can be distinguished
func (e *cRError) WithSubCode(code interface{}) IError {
	e.subCode = fmt.Sprintf("%v", code)

	return e
}

// WithField sets the error logData
func (e *cRError) WithField(key string, value interface{}) IError {
	e.logData[key] = value

	return e
}

func (e *cRError) WithInternalMessage(msg string) IError {
	return e.WithField("message", msg)
}

// InternalErrorDetails will give the details about the internal error
// for the current request
func (e *cRError) InternalErrorDetails() map[string]interface{} {
	ierr := map[string]interface{}{
		Code:    e.ErrorCode(),
		SubCode: e.SubCode(),
		Message: e.Message(),
	}

	if errorData := e.ErrorData(); errorData != nil {
		ierr[Errors] = errorData
	}

	return ierr
}

// PublicErrorDetails will give the details about the public error
// for the current request
func (e *cRError) PublicErrorDetails() map[string]interface{} {
	if e.publicErrorDetails.Code() == "" {
		return e.InternalErrorDetails()
	}
	return map[string]interface{}{
		Code:         e.publicErrorDetails.Code(),
		Message:      e.publicErrorDetails.Message(),
		DebugMessage: e.Message(),
	}
}

func (e *cRError) ErrorData() []map[string]interface{} {
	return e.errorData
}

// IsSubCode will check if the given sub code is same as sub code of error
func (e *cRError) IsSubCode(code string) bool {
	return e.subCode == code
}

// Log will log the error registered
func (e *cRError) Log() IError {
	data := e.logData
	data[InternalErrorDetails] = e.InternalErrorDetails()
	data[PublicErrorDetails] = e.PublicErrorDetails()

	log := logger.Get(e.ctx).WithError(e.GetInternalError()).WithFields(data)

	switch e.Level() {
	case LevelError:
		log.Error(e.PublicMessage())

	case LevelInfo:
		log.Info(e.PublicMessage())

	case LevelWarn:
		log.Warn(e.PublicMessage())

	case LevelFatal:
		log.Fatal(e.PublicMessage())
	}

	return e
}

func (e *cRError) GetTraceData() map[string]interface{} {
	data := e.logData
	data[InternalErrorDetails] = e.InternalErrorDetails()
	data[PublicErrorDetails] = e.PublicErrorDetails()

	return data
}
