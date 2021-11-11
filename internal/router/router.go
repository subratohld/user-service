package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subratohld/user-service/internal/handler"
	"github.com/subratohld/user-service/internal/service"
)

func Routes(userSvc service.User) http.Handler {
	router := mux.NewRouter()
	router.Path("/api/user/svc/health").Methods("GET").Handler(handler.HealthCheck{})
	router.Path("/api/user/svc/stats").Methods("GET").Handler(handler.Stats{})
	router.Path("/api/users").Methods("POST").Handler(handler.NewCreateUserHandler(userSvc))
	return router
}
