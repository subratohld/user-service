package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	xconfig "github.com/subratohld/config"
	xlogger "github.com/subratohld/logger"
	"go.uber.org/zap"
)

func New() xconfig.Reader {
	configReader, err := xconfig.New("/etc/config/config.yaml")
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
