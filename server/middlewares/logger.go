package middlewares

import (
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var disabled = []string{"/health"}

func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if slices.Contains(disabled, ctx.FullPath()) {
			return
		}

		var (
			message = "[REQUEST]"
			log     = logger.Info
			startT  = time.Now()
			fields  = []zap.Field{
				zap.String("method", ctx.Request.Method),
				zap.String("path", ctx.Request.URL.Path),
				zap.Time("begins", startT),
			}
		)

		log(message, fields...)

		ctx.Next()

		var (
			status = ctx.Writer.Status()
			takes  = time.Since(startT).Milliseconds()
			ends   = time.Now()
		)

		fields = append(fields,
			zap.Time("ends", ends),
			zap.Int("status", status),
			zap.Int64("takes (ms)", takes),
		)

		message = "[SUCCESS]"

		if status >= 400 {
			log = logger.Error
			message = "[FAIL]"
		}

		log(message, fields...)
	}
}
