package middleware

import (
	"github.com/gin-gonic/gin"
	"secret-server/app/metric"
	"strconv"
)

func Instrument() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := c.HandlerName()
		metric.TotalRequestCount.WithLabelValues(handler).Inc()

		c.Next()
		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method
		metric.RequestCount.WithLabelValues(status, method, handler).Inc()
	}
}
