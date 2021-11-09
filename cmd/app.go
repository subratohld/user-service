package main

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/subratohld/user-service/internal/api"
	"github.com/subratohld/user-service/internal/repository"
	"github.com/subratohld/user-service/internal/service"
	"go.uber.org/zap"
)

var rootCmd *cobra.Command = &cobra.Command{}

func main() {
	configReader := service.NewConfig()
	logger := service.NewLogger(configReader)
	db := service.DB(configReader)

	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo, logger)

	apis := api.APIs(userService)

	server := http.Server{
		Addr:    ":" + configReader.GetString("server.port"),
		Handler: apis,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal("Cannot start server", zap.Error(err))
	}
}

func startApp() {

}
