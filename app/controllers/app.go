package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"secret-server/app/constants"
	"secret-server/app/models"
	"secret-server/app/routing/response"
)

type appController struct {
	baseController
}

var (
	// App will hold the instance of appController
	App appController
	// AppRunningStatus : application running status
	AppRunningStatus = true
)

// NewAppController will initialize the new application controller to handle the method defined
func NewAppController() {
	App = appController{}
}

// @Summary Welcome
// @Description Welcome will return the welcome message as response with status code as 200
// @Tags System status
// @Success 200 {object} models.Base
// @Produce  json
// @Router / [get]
func (controller appController) Welcome(ctx *gin.Context) response.Response {
	responseContent := map[string]interface{}{
		"Message": "Welcome to Router :)",
	}
	responseStruct := response.NewResponse(ctx).
		SetResponse(responseContent).
		SetError(nil).
		SetStatusCode(http.StatusOK)
	return responseStruct
}

// @Summary Get Status
// @Description Status will provide the status of the service we been used in the application
// @Tags System status
// @Success 200 {object} models.Base
// @Produce  json
// @Router /status [get]
func (controller appController) Status(ctx *gin.Context) response.Response {
	resp := models.Base{}
	if AppRunningStatus {
		resp.Success = true
		return response.NewResponse(ctx).
			SetResponse(resp).
			SetError(nil).
			SetStatusCode(http.StatusOK)
	} else {
		resp.Success = false
		return response.NewResponse(ctx).
			SetResponse(resp).
			SetError(nil).
			SetStatusCode(http.StatusServiceUnavailable)
	}
}

// @Summary Ping test
// @Description Ping test, return the commit ID and container ID
// @Tags System status
// @Success 200 {object} models.Base
// @Produce  json
// @Router /ping [get]
func (controller appController) Ping(ctx *gin.Context) response.Response {
	result := make(map[string]interface{})

	result[constants.CommitID] = os.Getenv(constants.GitCommitHash)
	result[constants.ContainerID] = os.Getenv(constants.Hostname)

	responseStruct := response.NewResponse(ctx).
		SetResponse(result).
		SetError(nil).
		SetStatusCode(http.StatusOK)
	return responseStruct

}
