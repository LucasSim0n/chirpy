package config

import (
	"fmt"
	"net/http"
)

func (a *ApiConfig) MetricsHandler(w http.ResponseWriter, req *http.Request) {

	metrics := fmt.Sprintf(metricsTemplate, a.FileserverHits.Load())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(metrics))
}

func (a *ApiConfig) MetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			a.FileserverHits.Add(1)
			next.ServeHTTP(w, r)
		},
	)
}

const metricsTemplate = `
	<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>
	`
