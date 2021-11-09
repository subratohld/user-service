package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel string

const (
	DEBUG LogLevel = "debug"
	INFO  LogLevel = "info"
	WARN  LogLevel = "warn"
	ERROR LogLevel = "error"
)

type Encoding string

const (
	CONSOLE Encoding = "console"
	JSON    Encoding = "json"
)

type Logger interface {
	Disable()
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

type Params struct {
	Paths          []string
	Level          LogLevel
	Encoding       Encoding
	Verbose        bool
	ShowLogger     bool
	ShowCaller     bool
	ShowStacktrace bool
}

func New(param *Params) (Logger, error) {
	loggerName := ""
	callerKey := ""
	stackTraceKey := ""

	if param.ShowLogger {
		loggerName = "logger"
	}

	if param.ShowCaller {
		callerKey = "caller"
	}

	if param.ShowStacktrace {
		stackTraceKey = "stackTrace"
	}

	// If verbose is true then it will write logs on console only
	if param.Verbose {
		param.Paths = nil
		param.Encoding = CONSOLE
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        loggerName,
		CallerKey:      callerKey,
		StacktraceKey:  stackTraceKey,
		MessageKey:     "msg",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if param.Encoding == CONSOLE {
		encoderCfg.TimeKey = ""
		encoderCfg.StacktraceKey = ""
		encoderCfg.CallerKey = ""
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg := zap.Config{
		Development:   false,
		Encoding:      string(param.Encoding),
		EncoderConfig: encoderCfg,
	}

	level := zap.NewAtomicLevel()

	if param.Level == DEBUG {
		level.SetLevel(zap.DebugLevel)
	} else if param.Level == INFO {
		level.SetLevel(zap.InfoLevel)
	} else if param.Level == WARN {
		level.SetLevel(zap.WarnLevel)
	} else if param.Level == ERROR {
		level.SetLevel(zap.ErrorLevel)
	} else {
		level.SetLevel(zap.FatalLevel)
	}

	cfg.Level = level
	cfg.OutputPaths = param.Paths

	for _, path := range param.Paths {
		if path == "stdout" || path == "stderr" {
			continue
		}

		if _, err := os.OpenFile(path, os.O_CREATE, 0664); err != nil {
			fmt.Println("Failed to create log file: ", path, " due to error: ", err)
		}
	}

	lgr, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	zap.RedirectStdLogAt(lgr, zap.DebugLevel)

	lgr = lgr.WithOptions(zap.AddCallerSkip(1))

	return &logger{
		logger: lgr,
	}, err
}

func NewLoggerWithoutConfig() (Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.DisableCaller = true
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	config.EncoderConfig.TimeKey = ""
	config.EncoderConfig.StacktraceKey = ""
	config.EncoderConfig.CallerKey = ""
	config.Encoding = string(CONSOLE)
	lgr, err := config.Build()

	return &logger{logger: lgr}, err
}

type logger struct {
	logger *zap.Logger
}

func (l *logger) Disable() {
	l.logger = zap.NewNop()
}

func (l *logger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}
