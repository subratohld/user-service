package service

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/subratohld/config"
	xlogger "github.com/subratohld/logger"
	"go.uber.org/zap"
)

func NewConfig() config.Reader {
	configFile := os.Getenv("CONFIG_FILE")
	configReader, err := config.New(configFile)
	if err != nil {
		// This logger will write logs on console
		logger, _ := xlogger.NewLoggerWithoutConfig()
		logger.Error("Cannot read config file", zap.Error(err))
		os.Exit(0)
	}

	configReader.OnConfigChange(func(changes fsnotify.Event) {
		fmt.Println(changes)
	})
	configReader.WatchConfig()

	return configReader
}
