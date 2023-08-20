package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type LogStruct struct {
	processTime time.Duration
	URI         string
}

func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		str := LogStruct{
			processTime: time.Since(time.Now()),
			URI:         gctx.Request.URL.Path,
		}

		logger.Info(fmt.Sprintf("Hello Path: %v", str))
		gctx.Next()
	}
}
