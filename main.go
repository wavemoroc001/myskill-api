package main

import (
	"log"
	"myskill-api/db"
	"myskill-api/handler"
	"myskill-api/logger"
)

func main() {
	logger, cleanupLoggerFunc := logger.NewZap()
	defer cleanupLoggerFunc()

	database, cleanupDBConnFunc := db.NewDBConn()
	defer cleanupDBConnFunc()
	if database != nil {
		logger.Info("connect success")
	}
	//coll := skill.NewCollection(database)
	//page, _ := coll.FindAllByPageAndSort(context.Background(), 1, 2)
	//log.Printf("%#v\n%#v\n%#v\n%#v", page.TotalPage, page.TotalCount, page.PageSize, page.PageNo)
	engine := NewGinEngine()
	router := handler.NewRouter(engine, logger)
	handler.Register(router)
	srv, cleanupSRVFunc := NewHTTPServer(router)
	defer cleanupSRVFunc()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
