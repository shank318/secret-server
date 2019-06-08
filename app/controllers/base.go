package controllers

import (
	"github.com/gin-gonic/gin"
	"secret-server/app/crerrors"
	"secret-server/app/models"
)

type baseController struct {
}

func (b baseController) getInput(ctx *gin.Context, input models.IModel) (models.IModel, crerrors.IError) {
	if err := ctx.BindJSON(&input); err != nil {
		return nil, crerrors.NewCrError(ctx, crerrors.CodeBindingFailed, err)
	}
	err := input.Validate()
	if err != nil {
		return nil, crerrors.NewCrError(ctx, crerrors.CodeValidationError, err)
	}
	return input, nil
}
