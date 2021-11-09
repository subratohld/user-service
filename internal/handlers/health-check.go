package handlers

import "net/http"

type HealthCheck struct{}

func (hc HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service is up & running!"))
	w.WriteHeader(http.StatusOK)
}
