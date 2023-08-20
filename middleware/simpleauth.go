package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

// SimpleAuth is a middleware that authenticates the request using a simple
func SimpleAuth() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Set("timestamp", time.Now())
		gctx.Next()
	}
}
