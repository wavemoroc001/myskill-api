package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewGinEngine() *gin.Engine {
	engine := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type", "TransactionID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	engine.Use(cors.New(config))
	return engine
}

func NewHTTPServer(engine *Router) (*http.Server, func()) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	return srv, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			msg := fmt.Errorf("failed to shutdown http server: %w", err).Error()
			engine.logger.Fatal(msg)
		}
	}
}
