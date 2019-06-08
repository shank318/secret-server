package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	conf := GetConfig()
	name := reflect.TypeOf(conf).Name()
	assert.Equal(t, "AppConfig", name)
}

func TestLoadConfig(t *testing.T) {
	dir, _ := os.Getwd()
	LoadConfig(dir+"/../..", "sample")
	conf := GetConfig()
	assert.Equal(t, conf.Application.Mode, "debug")
	assert.Equal(t, conf.AuthUser.API.UserName, "api_user")
}
