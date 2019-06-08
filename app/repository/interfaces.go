package repository

import (
	"github.com/gin-gonic/gin"
	"secret-server/app/crerrors"
)

type IBaseRepo interface {
	FindByPkIdString(*gin.Context, interface{}, string) crerrors.IError
	Create(*gin.Context, interface{}) crerrors.IError
	Update(*gin.Context, interface{}, map[string]interface{}) crerrors.IError
}

type ISecretsRepo interface {
	IBaseRepo
}
