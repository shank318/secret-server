package repository

import (
	"secret-server/app/crerrors"
	"secret-server/app/providers/database"
	"secret-server/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type baseRepo struct{}

func (*baseRepo) FindByPkIdString(ctx *gin.Context, model interface{}, primaryKeyId string) crerrors.IError {
	table := utils.GetTypeName(model)
	err := getDbInstance().First(model, "id = ?", primaryKeyId).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":          table,
		"entity":         model,
		"primary_key_id": primaryKeyId,
	})
}

func (*baseRepo) Create(ctx *gin.Context, model interface{}) crerrors.IError {
	table := utils.GetTypeName(model)
	err := getDbInstance().Create(model).Error

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":  table,
		"entity": model,
	})
}


func (*baseRepo) Update(ctx *gin.Context, model interface{}, values map[string]interface{}) crerrors.IError {
	var err error
	table := utils.GetTypeName(model)
	// If the model is empty, it'll end up updating the whole table.
	if utils.IsEmpty(model) {
		err = utils.GetError("Empty model reference while update")
	} else {
		err = getDbInstance().Model(model).UpdateColumns(values).Error
	}

	return getErrorResponse(ctx, err, map[string]interface{}{
		"table":     table,
		"model":     model,
		"values":    values,
		"condition": model,
	})
}

func getDbInstance() *gorm.DB {
	return database.GetClient().Instance()
}

func getErrorResponse(ctx *gin.Context, err error, data map[string]interface{}) crerrors.IError {
	if err != nil {
		message := err.Error()
		errorCode := crerrors.CodeDatabaseError
		if gorm.IsRecordNotFoundError(err) {
			errorCode = crerrors.CodeRecordNotFound
		} else if strings.Contains(err.Error(), "Duplicate entry") {
			errorCode = crerrors.CodeRecordDuplicateEntry
		}
		return crerrors.NewCrError(ctx, errorCode, err).
			WithInternalMessage(message).
			WithFields(data).
			Log()
	}

	return nil
}

var BaseRepo IBaseRepo

func init() {
	BaseRepo = new(baseRepo)
}
