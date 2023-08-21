package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Database struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func NewDBConn(config Database) (*mongo.Database, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	auth := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}
	option := options.Client().ApplyURI(config.Host).SetAuth(auth)

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
