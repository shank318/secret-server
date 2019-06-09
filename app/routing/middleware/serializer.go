package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"secret-server/app/constants"
	"secret-server/app/routing/response"
)

// Serialize middleware formats the response in json format
func SerializeResponse(ctx *gin.Context) {
	ctx.Next()
	setResponse(ctx)
}

// sets the response and status code for the current request in json format
// data will be taken form the context value set by controller
func setResponse(ctx *gin.Context) {
	responseBody, status := prepareResponse(ctx)
	_, ok := responseBody.(map[string]interface{})
	ct := ctx.Request.Header.Get("accept")
	if (ct == "text/xml" || ct == "application/xml") && !ok {
		ctx.XML(status, responseBody)
		return
	}
	ctx.JSON(status, responseBody)
}

// if not available give the empty map
func prepareResponse(ctx *gin.Context) (interface{}, int) {
	data, exists := ctx.Get(constants.Response)

	if !exists {
		return map[string]interface{}{}, http.StatusOK
	}

	responseStruct := data.(response.Response)
	setHeaders(ctx, responseStruct.Headers())
	responseStruct.Log()
	return responseStruct.GetResponseBody(), responseStruct.StatusCode()
}

// setHeaders: will set the response header
func setHeaders(ctx *gin.Context, headers map[string]string) {
	for key, val := range headers {
		ctx.Header(key, val)
	}
}
