package logger

import (
	"context"
	"github.com/amku0xb/go-starter/config"
	"github.com/amku0xb/go-starter/constant"
	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger = logrus.New()
var logfile *os.File

func InitializeLogger(config *config.AppConfig) {

	l, e := logrus.ParseLevel(config.LogLevel)
	if e != nil {
		l = logrus.InfoLevel
	}

	logPath := config.LogFilePath
	var logFormatter = logrus.Logger{}
	if config.LogFormatJSON {
		logFormatter.Formatter = new(logrus.JSONFormatter)
	} else {
		logFormatter.Formatter = new(logrus.TextFormatter)
	}
	if logPath != "" {
		logfile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		if err == nil {
			*Logger = logrus.Logger{
				Out:       logfile,
				Formatter: logFormatter.Formatter,
				Hooks:     make(logrus.LevelHooks),
				Level:     l}
		}
		Logger.Info(constant.InitLogger)
	}
}

func WatchLoggerChanges(watcher *fsnotify.Watcher, config *config.AppConfig) {

	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				if ev.Name == config.LogFilePath && (ev.Op.String() == "REMOVE" || ev.Op.String() == "RENAME") {
					logfile.Close()
					f, err := os.OpenFile(config.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
					if err == nil {
						Logger.Out = f
						logfile = f
					} else {
						panic(err)
					}
					err = watcher.Add(logfile.Name())
					if err != nil {
						panic(err)
					}
				}
			case err := <-watcher.Errors:
				Logger.Error("error:", err)
			}
		}
	}()
	err := watcher.Add(config.LogFilePath)
	if err != nil {
		panic(err)
	}
}

type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

func GetLogger(ctx context.Context) *StructuredLoggerEntry {
	logger := ctx.Value(middleware.LogEntryCtxKey)
	if logger != nil {
		logEntry := logger.(*StructuredLoggerEntry)
		logEntry.Logger = logEntry.Logger.WithField("@timestamp", time.Now().Format(time.RFC3339Nano))
		return logEntry
	}
	return &StructuredLoggerEntry{Logger: Logger}
}
