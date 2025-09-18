package config

import (
	"chirpy/internal/auth"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (a *ApiConfig) UpdateRedHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error getting apikey from request: %s", err)
		return
	}
	if apiKey != a.PolkaKey {
		respondWithError(w, 401, "Unauthorized")
		log.Println("Provided apikey does not match.")
		return
	}
	data := upgrade{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(204)
		return
	}
	if data.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	userId, err := uuid.Parse(data.Data.UserId)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	_, err = a.DB.UpdateUserToRed(context.Background(), userId)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(204)

}

type upgrade struct {
	Event string  `json:"event"`
	Data  upgData `json:"data"`
}

type upgData struct {
	UserId string `json:"user_id"`
}
