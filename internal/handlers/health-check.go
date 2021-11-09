package handlers

import "net/http"

type HealthCheck struct{}

func (hc HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := new(jsonResp)
	resp.Message("Service is up & running!").Write(w, http.StatusOK)
}
