package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"secret-server/app/models"
	"secret-server/app/routing/response"
)

type secretController struct {
	baseController
}

var (
	// App will hold the instance of appController
	SecretApp secretController
	// AppRunningStatus : application running status
	SecretAppRunningStatus = true
)

func NewSecretController() {
	SecretApp = secretController{}
}

// @Summary Add a new secret
// @Description Add a new secret
// @Tags secret
// @Success 200 {object} models.Secret
// @Produce  json
// @Router /secret [post]
func (controller secretController) CreateSecret(ctx *gin.Context) response.Response {

	responseContent := models.SecretResponse{}
	responseStruct := response.NewResponse(ctx).
		SetResponse(responseContent).
		SetError(nil).
		SetStatusCode(http.StatusOK)
	return responseStruct
}
