package service

import (
	"os"

	"github.com/subratohld/config"
	xlogger "github.com/subratohld/logger"
	"go.uber.org/zap"
)

func NewLogger(reader config.Reader) xlogger.Logger {
	params := xlogger.Params{
		Paths:          reader.GetStringSlice("log.paths"),
		Level:          xlogger.LogLevel(reader.GetString("log.level")),
		Encoding:       xlogger.Encoding(reader.GetString("log.encoding")),
		Verbose:        reader.GetBool("log.verbose"),
		ShowLogger:     reader.GetBool("log.showLogger"),
		ShowCaller:     reader.GetBool("log.showCaller"),
		ShowStacktrace: reader.GetBool("log.showStacktrace"),
	}

	logger, err := xlogger.New(&params)
	if err != nil {
		// This logger will write logs on console
		lgr, _ := xlogger.NewLoggerWithoutConfig()
		lgr.Error("cannot create logger", zap.Error(err))
		os.Exit(0)
	}

	return logger
}
