package database

import (
	"cps/app/config"
	"os"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {
	client = new(dbProvider)
	c := GetClient()
	assert.Equal(t, client, c)
}

func TestDbProvider_Instance(t *testing.T) {
	provider := &dbProvider{}
	gormDb := new(gorm.DB)
	provider.instance = gormDb

	instance := provider.Instance()
	assert.Equal(t, gormDb, instance)
}

func TestInitialize(t *testing.T) {
	path, _ := os.Getwd()

	path = strings.Replace(path, "app/providers/database", "", 1)

	if os.Getenv("DRONE") == "true" {
		config.LoadConfig(path, "drone", "")
	} else {
		config.LoadConfig(path, "dev", "")
	}

	Initialize()

	assert.NotEqual(t, nil, GetClient())
}

func TestGetDatabasePath(t *testing.T) {
	path, _ := os.Getwd()

	path = strings.Replace(path, "app/providers/database", "", 1)

	if os.Getenv("DRONE") == "true" {
		config.LoadConfig(path, "drone", "")
	} else {
		config.LoadConfig(path, "dev", "")
	}

	databasePath := GetDatabasePath(config.GetConfig().Database)

	assert.NotEmpty(t, databasePath)
}
