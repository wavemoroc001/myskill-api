package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewDBConn() (*mongo.Database, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//auth := options.Credential{
	//	Username: "arise",
	//	Password: "arise123",
	//}
	auth := options.Credential{
		Username: "arise",
		Password: "arise123",
	}
	option := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(auth)
	conn, err := mongo.Connect(ctx, option)
	if err != nil {
		log.Fatal(fmt.Errorf("can not connect to mongodb: %w", err))
	}
	//db := conn.Database("myskill")
	db := conn.Database("ariskill")

	return db, func() {
		defer cancel()
		conn.Disconnect(ctx)
	}

}
