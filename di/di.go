package di

import (
	"myskill-api/config"
	"myskill-api/datastore"
	"myskill-api/db"
	"myskill-api/handler"
	"myskill-api/logger"
	"myskill-api/service"
	"net/http"
)

func Wire() (*http.Server, func()) {
	cfg := config.LoadEnv()
	logger, cleanupLoggerFunc := logger.NewZap()
	database, cleanupDBConnFunc := db.NewDBConn(cfg.Database)
	skillDataStore := datastore.NewSkillCollection(database)
	skillService := service.NewSkillService(skillDataStore)
	skillHandler := handler.NewSkillHandler(skillService)

	handlerSet := handler.NewHandlerSet(skillHandler)

	engine := handler.NewGinEngine()
	router := handler.NewRouter(engine, logger, handlerSet)
	handler.Register(router)

	srv, cleanupSRVFunc := handler.NewHTTPServer(router)

	return srv, func() {
		cleanupLoggerFunc()
		cleanupDBConnFunc()
		cleanupSRVFunc()
	}

}
