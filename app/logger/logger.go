//Package logger provides initializes logger and external logging service hooks
package logger

import (
	"os"
	"secret-server/app/config"
	"secret-server/app/constants"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	L           = logrus.NewEntry(logrus.StandardLogger())
	serviceData map[string]interface{}
)

// WithLogger returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithLogger(ctx *gin.Context, logger *logrus.Entry) {
	ctx.Set(constants.LOGGER, logger)
}

func Get(ctx *gin.Context) *logrus.Entry {
	return L
}

func Info(ctx *gin.Context, traceCode string, traceData map[string]interface{}) {
	traceData = getTraceData(ctx, traceData)

	L.WithFields(traceData).Info(traceCode)
}

func Debug(ctx *gin.Context, traceCode string, traceData map[string]interface{}) {
	traceData = getTraceData(ctx, traceData)

	L.WithFields(traceData).Debug(traceCode)
}

//func RzpError(ctx *gin.Context, rzpError crerrors.IError) {
//	Error(ctx, rzpError.ErrorCode(), rzpError.GetInternalError(), rzpError.GetTraceData())
//}

func Error(ctx *gin.Context, traceCode string, errorObj error, traceData map[string]interface{}) {
	traceData = getTraceData(ctx, traceData)

	L.WithFields(traceData).WithError(errorObj).Error(traceCode)
}

func Warn(ctx *gin.Context, traceCode string, traceData map[string]interface{}) {
	traceData = getTraceData(ctx, traceData)

	L.WithFields(traceData).Warn(traceCode)
}

func Fatal(ctx *gin.Context, traceCode string, traceData map[string]interface{}) {
	traceData = getTraceData(ctx, traceData)

	L.WithFields(traceData).Fatal(traceCode)
}

func Panic(ctx *gin.Context, traceCode string, traceData map[string]interface{}) {
	traceData = getTraceData(ctx, traceData)

	L.WithFields(traceData).Panic(traceCode)
}

func getTraceData(ctx *gin.Context, traceData map[string]interface{}) map[string]interface{} {
	finalTraceData := map[string]interface{}{}
	finalTraceData[constants.Service] = getServiceData()
	if ctx != nil {
		requestData, exists := ctx.Get(constants.Request)
		if exists == true {
			finalTraceData[constants.Request] = requestData
		}
	}

	if traceData != nil {
		finalTraceData[constants.Context] = traceData
	}

	return finalTraceData
}

func updateLog(log *logrus.Logger) {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
	log.SetLevel(getLevel())
	//TODO: fix the hook function and enable below
	//log.Hooks.Add(getSentryHooks())
}

func getLevel() logrus.Level {
	debugMode := config.GetConfig().Application.Mode
	if debugMode == "debug" {
		return logrus.DebugLevel
	}
	return logrus.InfoLevel
}

func getServiceData() map[string]interface{} {
	if serviceData == nil {
		serviceData = map[string]interface{}{}
		serviceData[constants.Env] = os.Getenv(constants.AppMode)
		serviceData[constants.ContainerID] = os.Getenv(constants.Hostname)
		serviceData[constants.CommitID] = os.Getenv(constants.GitCommitHash)
	}
	return serviceData
}
