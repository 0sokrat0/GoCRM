package middleware

import (
	"time"

	"GoCRM/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("HTTP request",
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.String("clientIP", c.ClientIP()),
			zap.Duration("latency", latency),
		)
	}
}
