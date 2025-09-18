package config

import (
	"context"
	"net/http"
)

func (a *ApiConfig) ResetHandler(w http.ResponseWriter, r *http.Request) {

	a.DB.ResetDB(context.Background())

	a.FileserverHits.Swap(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
