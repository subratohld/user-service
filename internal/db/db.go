package db

import (
	"os"

	xconfig "github.com/subratohld/config"
	xlogger "github.com/subratohld/logger"
	"github.com/subratohld/sqldb"
	"github.com/subratohld/user-service/internal/constant"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

func DB(configReader xconfig.Reader) sqldb.Sql {
	param := sqldb.Params{
		Username: configReader.GetString(constant.KEY_MYSQL_USER),
		Password: configReader.GetString(constant.KEY_MYSQL_PASSWORD),
		Host:     configReader.GetString(constant.KEY_MYSQL_HOST),
		Port:     configReader.GetString(constant.KEY_MYSQL_PORT),
		Database: configReader.GetString(constant.KEY_MYSQL_DATABASE),
	}

	db, err := sqldb.DB(param)
	if err != nil {
		logger, _ := xlogger.NewLoggerWithoutConfig()
		logger.Error("Could not connect to database", zap.Errors("errors", multierr.Errors(err)))
		os.Exit(0)
	}
	return db
}
