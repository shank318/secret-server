package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"secret-server/app/constants"
	"secret-server/app/logger"
	"secret-server/app/utils"
	"secret-server/app/crerrors"

	"github.com/gin-gonic/gin"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

func SetRequestPayload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceData := getRequestDetails(ctx)

		traceData[constants.Mode] = ctx.GetString(constants.Mode)
		traceData[constants.RequestID] = utils.GenerateUUID()
		traceData[constants.TaskID] = getTaskId(ctx)
		ctx.Set(constants.Request, traceData)

		input := make(map[string]interface{})
		collectRequestBody(ctx, input)

		collectQueryParam(ctx, input)
		collectPathParam(ctx, input)

		if ctx != nil && !utils.Contains([]string{"/", "/status"}, ctx.Request.URL.Path) {
			logger.Info(ctx, logger.TraceAppRequestData, input)
		}

		ctx.Next()
	}
}

func getRequestDetails(ctx *gin.Context) map[string]interface{} {
	data := map[string]interface{}{}

	if ctx.Request == nil {
		return data
	}

	data["uri"] = ctx.Request.RequestURI
	data["host"] = ctx.Request.Host
	data["referer"] = ctx.Request.Referer()

	return data
}

func getTaskId(ctx *gin.Context) string {
	taskId := ctx.Request.Header.Get(constants.X_RAZORPAY_TASK_ID)

	if len(taskId) == 0 {
		taskId = utils.GenerateUUID()
	}
	return taskId
}

// Collects the request body of request and sets as a map which can be accessed by context
// Request body will be parsed only when the request method is not GET
func collectRequestBody(ctx *gin.Context, input map[string]interface{}) crerrors.IError {
	if ctx.Request.Method == http.MethodGet {
		return nil
	}

	body, err := ctx.GetRawData()

	if err != nil {
		return crerrors.NewCrError(ctx, crerrors.CodeRuntimeError, err).Log()
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &input)
		read := ioutil.NopCloser(bytes.NewBuffer(body))

		ctx.Request.Body = read

		if err != nil {
			return crerrors.NewCrError(ctx, crerrors.CodeRuntimeError, err).Log()
		}
	}

	return nil
}

// Collects the query params available in the request URL
// based on the available data input map will be updated
func collectQueryParam(ctx *gin.Context, input map[string]interface{}) {
	values := ctx.Request.URL.Query()

	for k, v := range values {
		//
		// v is always an array
		// If there's only one value, we want to directly assign
		// it to the key. If there's more than one value, we want
		// to assign the whole array as the value for the key.
		//
		if len(v) == 1 {
			input[k] = v[0]
		} else {
			input[k] = v
		}
	}
}

// Collects the route params available in the request URL
// based on the available data input map will be updated
func collectPathParam(ctx *gin.Context, input map[string]interface{}) {
	params := ctx.Params

	for _, param := range params {
		input[param.Key] = param.Value
	}
}
