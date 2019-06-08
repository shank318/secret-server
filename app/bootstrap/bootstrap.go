// Package bootstrap - initialize all the components required for the application to start
// This will set the application environments and sets default logging for framework
package bootstrap

import (
	"io/ioutil"
	"secret-server/app/config"
	"secret-server/app/controllers"
	"secret-server/app/logger"
	"secret-server/app/providers/database"
	"secret-server/app/routing/router"
	"secret-server/docs"

	"github.com/gin-gonic/gin"
)

var initializeRouter = router.Initialize
var initializeLogger = logger.InitLogrus
var initializeDatabase = database.Initialize
var loadConfig = config.LoadConfig
var newAppController = controllers.NewAppController
var newSecretController = controllers.NewSecretController
var setGinMode = gin.SetMode

// Initialize : initializes all required application components
func Initialize(basePath string, env string) {
	baseInit(basePath, env)
	swaggerInfo()
	initializeRouter()
}

// BaseInit : Basic initializations required for tests
func baseInit(basePath string, env string) {
	loadConfig(basePath, env)

	setEnvironment()
	initProviders(basePath)
	initializeRequestHandlers()
}

// initProviders : Provider initialization is done here
// There initiated providers will be available across the application
func initProviders(basePath string) {
	initializeDatabase()
	initializeLogger()
}

// initializeRequestHandlers : initializing request handlers
func initializeRequestHandlers() {
	newAppController()
}

// setEnvironment : sets application gin environment based on application mode
// gin default log writer will be changed for the `release` mode
func setEnvironment() {
	appMode := config.GetConfig().Application.Mode
	setGinMode(appMode)

	if appMode == gin.ReleaseMode {
		// Disabling gin logs for release mode
		gin.DefaultWriter = ioutil.Discard
	}
}

func swaggerInfo() {
	// Programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger API documentation for Secret Server"
	docs.SwaggerInfo.Description = "Secret Server"
	docs.SwaggerInfo.Version = "1.0"
}
