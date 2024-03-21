package middleware


import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware returns a Gin middleware handler function for logging.
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// End timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		logger.Info("request",
			zap.String("method", method),
			zap.Int("status", statusCode),
			zap.String("path", path),
			zap.String("ip", clientIP),
			zap.Duration("latency", latency),
		)
	}
}


