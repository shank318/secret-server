package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"secret-server/app/constants"
	"secret-server/app/crerrors"
	"secret-server/app/metric"
	"secret-server/app/models"
	"secret-server/app/repository"
	"secret-server/app/routing/response"
	"secret-server/app/service"
	"secret-server/app/utils"
	"strconv"
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
// @Success 200 {object} models.SecretResponse
// @Param secret formData string true "This text will be saved as a secret"
// @param expireAfterViews formData int true "The secret won't be available after the given number of views. It must be greater than 0."
// @Param expireAfter formData int true "The secret won't be available after the given time. The value is provided in minutes. 0 means never expires"
// @Accept x-www-form-urlencoded
// @Produce  json,xml
// @Router /secret [post]
func (controller secretController) CreateSecret(ctx *gin.Context) response.Response {
	defer utils.TraceExecutionTime(ctx, constants.CreateSecret, utils.GetCurrentTimeStampInMiliSec())
	defer metric.RequestCount.WithLabelValues(constants.CreateSecret).Inc()
	responseStruct := response.
		NewResponse(ctx)

	var ierr crerrors.IError
	ctx.Request.ParseForm()
	secretText := ctx.PostForm("secret")
	expireAfterViews, err := strconv.Atoi(ctx.PostForm("expireAfterViews"))
	if err != nil {
		ierr = crerrors.NewCrError(ctx, crerrors.BadRequestErrorCode, err)
	}
	expireAfter, err := strconv.Atoi(ctx.PostForm("expireAfter"))
	if err != nil {
		ierr = crerrors.NewCrError(ctx, crerrors.BadRequestErrorCode, err)
	}
	if ierr != nil {
		responseStruct.SetError(ierr)
		return responseStruct
	}
	secretReq := &models.SecretRequest{
		SecretText:       secretText,
		ExpireAfterViews: expireAfterViews,
		ExpireAfter:      expireAfter,
	}
	err = secretReq.Validate()
	if err != nil {
		responseStruct.SetError(crerrors.NewCrError(ctx, crerrors.BadRequestErrorCode, err))
		return responseStruct
	}

	secretService := service.NewSecretService(repository.SecretsRepo)
	secret, iError := secretService.CreateSecret(ctx, secretReq)
	if iError != nil {
		responseStruct.SetError(iError)
		return responseStruct
	}
	responseContent := models.SecretResponse{
		Hash:           secret.ID,
		SecretText:     secret.SecretText,
		CreatedAt:      secret.CreatedAt,
		ExpiresAt:      secret.ExpiresAt,
		RemainingViews: secret.RemViews,
	}
	responseStruct.SetResponse(responseContent).
		SetStatusCode(http.StatusOK)
	return responseStruct
}

// @Summary Get a secret by ID
// @Description Get a secret by ID
// @Tags secret
// @Success 200 {object} models.SecretResponse
// @Param hash path string true "Unique hash to identify the secret"
// @Produce  json,xml
// @Router /secret/{hash} [get]
func (controller secretController) GetSecret(ctx *gin.Context) response.Response {
	defer utils.TraceExecutionTime(ctx, constants.GetSecret, utils.GetCurrentTimeStampInMiliSec())
	defer metric.RequestCount.WithLabelValues(constants.GetSecret).Inc()
	responseStruct := response.
		NewResponse(ctx)

	hash := ctx.Param("hash")
	secretService := service.NewSecretService(repository.SecretsRepo)
	secret, iError := secretService.GetSecret(ctx, hash)
	if iError != nil {
		responseStruct.SetError(iError)
		return responseStruct
	}
	responseContent := models.SecretResponse{
		Hash:           secret.ID,
		SecretText:     secret.SecretText,
		CreatedAt:      secret.CreatedAt,
		ExpiresAt:      secret.ExpiresAt,
		RemainingViews: secret.RemViews,
	}
	responseStruct.SetResponse(responseContent).
		SetStatusCode(http.StatusOK)
	return responseStruct
}
