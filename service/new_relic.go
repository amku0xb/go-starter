package service

import (
	"context"
	"github.com/amku0xb/go-starter/config"
	"github.com/amku0xb/go-starter/constant"
	"github.com/amku0xb/go-starter/logger"
	"github.com/newrelic/go-agent"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var app newrelic.Application

func InitNewRelic(config *config.AppConfig) newrelic.Application {
	nrConfig := newrelic.NewConfig(config.NewRelic["app_name"].(string), config.NewRelic["license_key"].(string))
	nrApp, err := newrelic.NewApplication(nrConfig)
	log := logger.GetLogger(context.Background())
	if err == nil {
		log.Logger.Info(constant.InitNewRelic)
	} else {
		log.Logger.WithError(err).Error(constant.InitNewRelicFail)
	}
	app = nrApp
	return nrApp
}

func getApp() newrelic.Application {
	return app
}

func RecordCustomMetric(metric string, value float64) {
	if getApp() == nil {
		logrus.Info(constant.NewRelicAgentNotFound)
		return
	}
	getApp().RecordCustomMetric(metric, value)
}

func RecordHttpTransaction(req *http.Request, name string) (*http.Response, error) {
	client := new(http.Client)
	client.Timeout = time.Duration(10 * time.Second)

	if getApp() == nil {
		return client.Do(req)
	}
	txn := getApp().StartTransaction(name, nil, req)
	defer txn.End()
	response, err := client.Do(req)
	return response, err
}
