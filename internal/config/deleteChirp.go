package config

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (a *ApiConfig) DeleteChirpHandler(w http.ResponseWriter, r *http.Request) {

	t, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error getting jwt: %s", err)
		return
	}

	userId, err := auth.ValidateJWT(t, a.Secret)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error validating jwt: %s", err)
		return
	}

	chirpId := r.PathValue("chirpID")

	parsedId, err := uuid.Parse(chirpId)
	if err != nil {
		respondWithError(w, 403, "Bad request")
		log.Printf("Error parsing chirp id: %s", err)
		return
	}

	deleteChirp := database.DeleteChirpByIdParams{
		ID:     parsedId,
		UserID: userId,
	}

	_, err = a.DB.DeleteChirpById(context.Background(), deleteChirp)
	if err == sql.ErrNoRows {
		respondWithError(w, 403, "Unauthorized")
	} else if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error deleting chirp from db: %s", err)
		return
	}

	w.WriteHeader(204)
}
