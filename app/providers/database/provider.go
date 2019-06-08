package database

import (
	"secret-server/app/logger"
	"secret-server/app/crerrors"
	"secret-server/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

// dbProvider will hold the connection related data
type dbProvider struct {
	instance           *gorm.DB
	callbacks          []func(*gorm.DB)
	path               string
	dialect            string
	maxOpenConnections int
	maxIdleConnections int
	connMaxLifetime    time.Duration
}

// Instance returns an instance of the database
func (dbProvider *dbProvider) Instance() *gorm.DB {
	return dbProvider.instance
}

// Ping checks the connectivity to the database server
func (dbProvider *dbProvider) Ping() error {
	// dbProvider.instance.DB().Ping()
	// Ping only checks if the connection is available in the pool
	// If not and connection limit is not reached it'll create one connection
	// If the instance of connection is available in the pool and if it was killed by the database server
	// 	- In this case Ping will return the available instance as it doesn't know the connection was closed by server
	//	- To handle this we execute a query, by doing which the connection will re-established with the server

	var err error

	if _, err = dbProvider.instance.DB().Exec("SELECT 1"); err != nil {
		logger.Error(nil, crerrors.CodeDatabaseError, err, nil)
	}

	return err
}

// Connects to the database specified the provider
func (dbProvider *dbProvider) connect() error {
	var err error

	dbProvider.instance, err = gorm.Open(dbProvider.dialect, dbProvider.path)

	if err != nil {
		logger.Error(nil, crerrors.CodeDatabaseError, err, nil)
	}

	if gin.IsDebugging() {
		dbProvider.instance.LogMode(true)
	}

	dbProvider.instance.DB().SetMaxIdleConns(dbProvider.maxIdleConnections)
	dbProvider.instance.DB().SetMaxOpenConns(dbProvider.maxOpenConnections)
	dbProvider.instance.DB().SetConnMaxLifetime(dbProvider.connMaxLifetime * time.Second)
	return nil
}

// setCreatedTimeStamp will check if given scope has created_at column.
// If yes and its not empty it'll update the field with current unix timestamp
func setCreatedTimeStamp(scope *gorm.Scope) {
	if !scope.HasError() {
		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			// If its blank then only we update it with current timestamp
			// This is done because sometimes we may get the timestamp from
			// other service or we might need to set it ourselves
			if createdAtField.IsBlank {
				createdAtField.Set(utils.GetCurrentTimeStamp())
			}
		}

		setUpdatedTimeStamp(scope)
	}
}

// setUpdatedTimeStamp will check if given scope has updated_at column.
// If yes then it'll update the field with current unix timestamp
func setUpdatedTimeStamp(scope *gorm.Scope) {
	scope.SetColumn("UpdatedAt", utils.GetCurrentTimeStamp())
}
