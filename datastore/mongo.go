package datastore

import (
	"context"
	"github.com/amku0xb/go-starter/config"
	"github.com/amku0xb/go-starter/constant"
	"github.com/amku0xb/go-starter/logger"
	"github.com/mongodb/mongo-go-driver/mongo"
	"sync"
)

type MongoStore struct {
	db      *mongo.Database
	session *mongo.Client
}

func NewMongoStore(config *config.AppConfig) *MongoStore {
	var mongoStore *MongoStore
	db, session := connect(config)
	if db != nil && session != nil {
		mongoStore = new(MongoStore)
		mongoStore.db = db
		mongoStore.session = session
		return mongoStore
	}

	log := logger.GetLogger(context.Background())
	log.Logger.Fatalf(constant.InitMongoFail, config.DatabaseName)

	return nil
}

func connect(generalConfig *config.AppConfig) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(generalConfig)
	})

	return db, session
}

func connectToMongo(generalConfig *config.AppConfig) (a *mongo.Database, b *mongo.Client) {
	var err error
	log := logger.GetLogger(context.Background())

	session, err := mongo.NewClient(generalConfig.DatabaseHost)
	if err != nil {
		log.Logger.Fatal(err)
	}
	err = session.Connect(context.TODO())
	if err != nil {
		log.Logger.Fatal(err)
	}

	var DB = session.Database(generalConfig.DatabaseName)
	log.Logger.Info(constant.InitMongoSuccess, generalConfig.DatabaseName)

	return DB, session
}
