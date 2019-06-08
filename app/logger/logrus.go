package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// initLogrus initializes logrus and adds the sentry hook
func InitLogrus() {
	logLevel := logrus.DebugLevel

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	writers := make([]io.Writer, 0, 2)

	writers = append(writers, os.Stdout)

	logLevel = logrus.InfoLevel

	mw := io.MultiWriter(writers...)
	logrus.SetOutput(mw)
	logrus.SetLevel(logLevel)

	logrus.AddHook(StandardHook())
}
