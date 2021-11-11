package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/subratohld/config"
	"github.com/subratohld/logger"
	"github.com/subratohld/user-service/internal/repository"
	"github.com/subratohld/user-service/internal/router"
	"github.com/subratohld/user-service/internal/server"
	"github.com/subratohld/user-service/internal/service"
	"go.uber.org/zap"
)

func main() {
	app := new(App).init()

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// It will keep monitoring for any interruption
	moniterSignals := func() {
		sig := <-quit

		app.svr.SetKeepAlivesEnabled(false)
		app.stop(sig.String())

		close(done)
	}
	go moniterSignals()

	// It will start the server, if successfully started then it will be blocked here
	// because  http's ListenAndServe uses goroutine
	app.start()

	// The app will be blocked here as well untill moniterSignals goroutine closes 'done' channel
	<-done
	app.logger.Info("Server has been stop gracefully")
}

type App struct {
	configReader config.Reader
	logger       logger.Logger
	svr          server.Server
}

func (app *App) init() *App {
	configReader := service.NewConfig()
	logger := service.NewLogger(configReader)

	db := service.DB(configReader)
	userRepo := repository.NewUserRepo(db)
	userSvc := service.NewUserService(userRepo, logger)
	handler := router.Routes(userSvc)

	svr := server.New(configReader, handler)

	app.configReader = configReader
	app.logger = logger
	app.svr = svr
	return app
}

func (app *App) start() {
	app.logger.Info("Some info", zap.Int("uid", os.Getuid()), zap.Int("pid", os.Getpid()))
	port := app.configReader.GetString("server.port")
	app.logger.Info("starting server at port " + port)

	// It will start the server, it is blocking in nature
	if err := app.svr.ListenAndServe(); err != nil {
		app.logger.Fatal("can not start server gracefully", zap.Error(err))
	}

	app.logger.Info("server started, it is listening at port " + port)
}

func (app *App) stop(signalMsg string) {
	app.logger.Info("stoping server", zap.String("signal", signalMsg))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := app.svr.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		app.logger.Fatal("server could not stop gracefully", zap.Error(err))
	}
}
