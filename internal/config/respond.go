package config

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	resp := errorResponse{
		Error: message,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error marshaling json response: %s", err)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithJson(w http.ResponseWriter, code int, payload any) {

	data, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, 500, genericError)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
