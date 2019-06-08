package response

import (
	"encoding/json"
	"net/http"
	"secret-server/app/crerrors"
	"secret-server/app/logger"

	"github.com/gin-gonic/gin"
	"secret-server/app/constants"
)

// Response interface for response struct
type Response interface {
	SetBody(interface{}) Response
	SetError(crerrors.IError) Response
	SetResponse(interface{}) Response
	SetHeaders(map[string]string) Response
	SetStatusCode(int) Response
	SetResponseType(string) Response
	Log()
	StatusCode() int
	ResponseType() string
	Body() interface{}
	ErrorBody() interface{}
	Error() crerrors.IError
	Headers() map[string]string
	GetResponseBody() interface{}
	String() string
}

// response struct will manage the response which has to be sent
type response struct {
	Response
	err          crerrors.IError
	body         interface{}
	statusCode   int
	responseType string
	headers      map[string]string
	ctx          *gin.Context
}

// NewResponse will create a new response instance
func NewResponse(ctx *gin.Context) Response {
	statusCode := http.StatusOK

	return &response{
		statusCode: statusCode,
		ctx:        ctx,
	}
}

// SetHeader will set the headers which has to be sent in the response
func (r *response) SetHeaders(headers map[string]string) Response {
	r.headers = headers

	return r
}

// AppendHeaders will append the given headers with the existing headers
func (r *response) AppendHeaders(headers map[string]string) Response {
	for key, val := range headers {
		r.headers[key] = val
	}

	return r
}

// Header will give back the headers which has to be set for the response
func (r *response) Headers() map[string]string {
	return r.headers
}

// SetStatusCode will set the code which will be displayed in with the error message
func (r *response) SetStatusCode(code int) Response {
	r.statusCode = code

	return r
}

// SetBody will set the response body. which will be sent as response if there are no rzperror
func (r *response) SetBody(data interface{}) Response {
	r.body = data

	return r
}

// SetResponseType will defined the response format
func (r *response) SetResponseType(responseType string) Response {
	r.responseType = responseType

	return r
}

// SetError will set the error attribute with the given error
func (r *response) SetError(ierr crerrors.IError) Response {
	if ierr != nil {
		r.statusCode = ierr.StatusCode()
	}

	r.err = ierr

	return r
}

// SetResponse will set the response attribute with the given response
func (r *response) SetResponse(response interface{}) Response {
	r.body = response

	return r
}

// StatusCode will give the status code which has to be sent as response
func (r *response) StatusCode() int {
	if r.err != nil && r.err.StatusCode() != 0 {
		return r.err.StatusCode()
	}

	return r.statusCode
}

// Body will provide the response body which was set earlier
func (r *response) Body() interface{} {
	return r.body
}

func (r *response) ErrorBody() interface{} {
	return r.err.PublicErrorDetails()
}

// ResponseType will give the response format which has to be followed for the current response
func (r *response) ResponseType() string {
	return r.responseType
}

// Error will give the error instance set
func (r *response) Error() crerrors.IError {
	return r.err
}

// GetResponseBody will provide the response body which has to set to the client
// If there is no error set then response body will be same as Body
// In case of error response will the formatted map of error
func (r *response) GetResponseBody() interface{} {
	if r.err != nil {
		return r.ErrorBody()
	}

	return r.Body()
}

// Log will log the response with headers which will be sent to the client
func (r *response) Log() {
	data := map[string]interface{}{
		constants.Response: r.GetResponseBody(),
		constants.Headers:  r.Headers(),
	}
	logger.Info(r.ctx, logger.TraceAppResponseData, data)
}

func (r *response) String() string {
	out, _ := json.Marshal(r)
	return string(out)
}
