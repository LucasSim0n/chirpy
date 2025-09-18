package config

import "net/http"

func HelthHandler(w http.ResponseWriter, req *http.Request) {

	ok := []byte("OK")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write(ok)
}
