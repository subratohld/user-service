package logger

import (
	"os"

	xconfig "github.com/subratohld/config"
	xlogger "github.com/subratohld/logger"
	"github.com/subratohld/user-service/internal/constant"
	"go.uber.org/zap"
)

func New(reader xconfig.Reader) xlogger.Logger {
	params := xlogger.Params{
		Paths:          reader.GetStringSlice(constant.KEY_LOG_PATHS),
		Level:          xlogger.LogLevel(reader.GetString(constant.KEY_LOG_LEVEL)),
		Encoding:       xlogger.Encoding(reader.GetString(constant.KEY_LOG_ENCODING)),
		Verbose:        reader.GetBool(constant.KEY_LOG_VERBOSE),
		ShowLogger:     reader.GetBool(constant.KEY_LOG_SHOWLOGGER),
		ShowCaller:     reader.GetBool(constant.KEY_LOG_SHOWCALLER),
		ShowStacktrace: reader.GetBool(constant.KEY_LOG_SHOWSTRACKTRACE),
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
