package app

import (
	"github.com/amku0xb/go-starter/config"
	"github.com/amku0xb/go-starter/constant"
	"github.com/amku0xb/go-starter/datastore"
	"github.com/amku0xb/go-starter/logger"
	"github.com/amku0xb/go-starter/service"
	"net/http"
)

type IApp interface {
	Init()
	Start(string) error
	GetConfig() *config.AppConfig
}

type App struct {
	config *config.AppConfig
}

func NewApp(configFilePath string) IApp {
	appConfig := config.LoadConfig(configFilePath)
	return &App{
		config: appConfig,
	}
}

func (a *App) GetConfig() *config.AppConfig {
	return a.config
}

func (a *App) Init() {
	logger.InitializeLogger(a.config)
	service.InitNewRelic(a.config)
	datastore.InitMgoStore(a.config)
	datastore.NewMongoStore(a.config)
	service.InitRedisClient(a.config)
}

func (a *App) Start(port string) error {
	logger.Logger.Infoln(constant.ServerStarted)
	return http.ListenAndServe(":"+port, nil)
}
