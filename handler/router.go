package handler

import (
	"github.com/gin-gonic/gin"
	"myskill-api/middleware"
	"net/http"
)

func Register(r *Router) {
	r.GET("/users", middleware.Logger(r.logger), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
