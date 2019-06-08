package bootstrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialize(t *testing.T) {
	oldInitializeRouter := initializeRouter
	oldInitializeLogger := initializeLogger
	oldLoadConfig := loadConfig
	oldNewAppController := newAppController
	oldSecretController := newSecretController
	oldSetGinMode := setGinMode

	defer func() { initializeRouter = oldInitializeRouter }()
	defer func() { initializeLogger = oldInitializeLogger }()
	defer func() { loadConfig = oldLoadConfig }()
	defer func() { newAppController = oldNewAppController }()
	defer func() { newSecretController = oldSecretController }()
	defer func() { setGinMode = oldSetGinMode }()

	initializeRouterCalled := false
	initializeLoggerCalled := false
	loadConfigCalled := false
	newAppControllerCalled := false
	newTerminalControllerCalled := false
	setGinModeCalled := false

	initializeRouter = func() {
		initializeRouterCalled = true
	}
	initializeLogger = func() {
		initializeLoggerCalled = true
	}
	loadConfig = func(basePath string, env string) {
		loadConfigCalled = true
	}
	newAppController = func() {
		newAppControllerCalled = true
	}
	newSecretController = func() {
		newTerminalControllerCalled = true
	}
	setGinMode = func(value string) {
		setGinModeCalled = true
	}

	Initialize(".", "")

	assert.Equal(t, true, initializeRouterCalled)
	assert.Equal(t, true, initializeLoggerCalled)
	assert.Equal(t, true, loadConfigCalled)
	assert.Equal(t, true, newAppControllerCalled)
	assert.Equal(t, true, newTerminalControllerCalled)
	assert.Equal(t, true, setGinModeCalled)

}
