package datastore

import (
	"context"
	"github.com/amku0xb/go-starter/config"
	"github.com/amku0xb/go-starter/constant"
	"github.com/amku0xb/go-starter/logger"
	"github.com/globalsign/mgo"
	"time"
)

type M map[string]interface{}

type MgoStore struct {
	Session *mgo.Session
}

func InitMgoStore(config *config.AppConfig) *MgoStore {
	log := logger.GetLogger(context.Background())

	dialInfo, err := mgo.ParseURL(config.DatabaseHost)
	if err != nil {
		log.Logger.WithError(err).Error(constant.MgoUrlParseErr)
		return nil
	}

	dialInfo.Database = config.DatabaseName
	dialInfo.PoolLimit = config.MgoPoolSize
	dialInfo.Timeout = time.Duration(config.MgoTimeOut) * time.Second

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Logger.WithError(err).Error(constant.MgoSesInitFail)
		return nil
	}

	session.SetMode(mgo.Monotonic, true)
	mgoStore := new(MgoStore)
	mgoStore.Session = session

	log.Logger.Info(constant.InitMgoDbSuccess)
	return mgoStore
}
