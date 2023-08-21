package main

import (
	"log"
	"myskill-api/di"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	srv, cleanupFunc := di.Wire()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		cleanupFunc()
	}()
}
