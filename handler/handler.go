package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	*gin.Engine
	logger *zap.Logger
	*Set
}

type Set struct {
	//userHandler  UserHandler
	SkillHandler SkillHandler
}

func NewHandlerSet(
	//userHandler UserHandler,
	skillHandler SkillHandler) *Set {
	return &Set{
		//userHandler:  userHandler,
		SkillHandler: skillHandler,
	}
}

func NewRouter(engine *gin.Engine, logger *zap.Logger, set *Set) *Router {
	return &Router{
		engine,
		logger,
		set,
	}
}
