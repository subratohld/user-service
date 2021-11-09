package api

import (
	"github.com/gorilla/mux"
	"github.com/subratohld/user-service/internal/handlers"
	"github.com/subratohld/user-service/internal/service"
)

func APIs(svc service.User) *mux.Router {
	router := mux.NewRouter()
	router.Path("/api/user/svc/health").Methods("GET").Handler(handlers.HealthCheck{})
	router.Path("/api/user/svc/stats").Methods("GET").Handler(handlers.Stats{})
	router.Path("/api/users").Methods("POST").Handler(handlers.CreateUser{Svc: svc})
	return router
}
