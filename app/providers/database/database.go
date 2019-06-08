// Package database will manage the connection to the database
package database

import (
	"secret-server/app/config"
)

// interface to be implemented in all child structs
// This should provide the connection path
type database interface {
	getDatabasePath() string
}

// Client holds the database connection
var (
	client *dbProvider
)

// Initialize will initialize the connection to the dialect
func Initialize() {
	if client == nil {
		client = getDbProvider()
		client.connect()
	}
}

// GetClient will give the data base client for the mode set in context
func GetClient() IDbProvider {
	return client
}

func GetDatabasePath(databaseConfig config.DatabaseConfig) string {
	return new(mysql).getDatabasePath(databaseConfig)
}

func getDbProvider() *dbProvider {
	databaseConfig := config.GetConfig().Database

	provider := &dbProvider{
		dialect:            databaseConfig.Dialect,
		maxIdleConnections: databaseConfig.MaxIdleConnections,
		maxOpenConnections: databaseConfig.MaxOpenConnections,
		connMaxLifetime:    databaseConfig.ConnectionMaxLifetime,
		path:               GetDatabasePath(databaseConfig),
	}

	return provider
}
