package main

import (
	"flag"
	app2 "github.com/amku0xb/go-starter/app"
	"github.com/amku0xb/go-starter/logger"
	"github.com/fsnotify/fsnotify"
)

func main() {
	var configFilePath, serverPort string
	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.StringVar(&serverPort, "server_port", "8000", "port on which server runs")
	flag.Parse()

	app := app2.NewApp(configFilePath)
	app.Init()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()
	if app.GetConfig().LogRotation {
		go logger.WatchLoggerChanges(watcher, app.GetConfig())
	}
	app.Start(serverPort)
}
