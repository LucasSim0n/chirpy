package config

import (
	"chirpy/internal/database"
	"context"
	"log"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (a *ApiConfig) GetChirpsHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Query().Get("author_id")
	var chirps []database.Chirp
	var err error
	if userId != "" {

		parsedId, err := uuid.Parse(userId)
		if err != nil {
			respondWithError(w, 500, genericError)
			log.Printf("Error parsing user id: %s", err)
			return
		}
		chirps, err = a.DB.GetChirpsByUserId(context.Background(), parsedId)
		if err != nil {
			respondWithError(w, 404, "Not found")
			log.Printf("Chirps not found: %s", err)
			return
		}

	} else {
		chirps, err = a.DB.GetAllChirps(context.Background())
		if err != nil {
			respondWithError(w, 500, genericError)
			log.Printf("Database error getting chirps: %s", err)
			return
		}
	}

	var parsedChirps []responseChirp

	for _, chirp := range chirps {
		pChirp := parseDBChirpToResponse(chirp)
		parsedChirps = append(parsedChirps, pChirp)
	}
	order := r.URL.Query().Get("sort")
	if order == "desc" {
		sort.Slice(parsedChirps, func(i, j int) bool { return parsedChirps[i].CreatedAt.After(parsedChirps[j].CreatedAt) })
	}
	respondWithJson(w, 200, parsedChirps)

}

func (a *ApiConfig) GetAChirpHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("chirpID")
	if id == "" {
		respondWithError(w, 404, "chirp not found")
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, 404, "chirp not found")
		log.Printf("id %s not parseable. error: %s", id, err)
		return
	}
	chirp, err := a.DB.GetChirpByID(context.Background(), parsedID)
	if err != nil {
		respondWithError(w, 404, "chirp not found")
		return
	}

	parsedChirp := parseDBChirpToResponse(chirp)

	respondWithJson(w, 200, parsedChirp)
}

func parseDBChirpToResponse(chirp database.Chirp) responseChirp {
	return responseChirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserId:    chirp.UserID,
	}
}
