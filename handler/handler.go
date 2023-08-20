package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	*gin.Engine
	logger *zap.Logger
}

func NewRouter(engine *gin.Engine, logger *zap.Logger) *Router {
	return &Router{
		engine,
		logger,
	}
}
