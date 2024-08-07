package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		log.Info().
			Str("method", c.Request.Method).
			Str("url", c.Request.URL.String()).
			Str("remote_addr", c.ClientIP()).
			Int("status", status).
			Dur("duration", duration).
			Str("user_agent", c.Request.UserAgent()).
			Msg("Request processed")
	}
}

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Logger = log.With().Caller().Logger()
}
