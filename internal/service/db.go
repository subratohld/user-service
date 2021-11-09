package service

import (
	"os"

	"github.com/subratohld/config"
	"github.com/subratohld/logger"
	"github.com/subratohld/sqldb"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

func DB(configReader config.Reader) sqldb.Sql {
	param := sqldb.Params{
		Username: configReader.GetString("db.mysql.username"),
		Password: configReader.GetString("db.mysql.password"),
		Host:     configReader.GetString("db.mysql.host"),
		Port:     configReader.GetString("db.mysql.port"),
		Database: configReader.GetString("db.mysql.database"),
	}
	db, err := sqldb.DB(param)
	if err != nil {
		logger, _ := logger.NewLoggerWithoutConfig()
		logger.Error("Could not connect to database", zap.Errors("errors", multierr.Errors(err)))
		os.Exit(0)
	}
	return db
}
