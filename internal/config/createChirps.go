package config

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	maxValidChirp = 140
	genericError  = "something went wrong"
)

type parameters struct {
	Body   string    `json:"body"`
	UserId uuid.UUID `json:"user_id"`
}

type responseChirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserId    uuid.UUID `json:"user_id"`
}

func (a *ApiConfig) ChirpHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error decoding request: %s", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error Getting Bearer token: %s", err)
		return
	}
	UserId, err := auth.ValidateJWT(token, a.Secret)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error validating jwt: %s", err)
	}

	if len(params.Body) > maxValidChirp {
		respondWithError(w, 400, "chirp too long")
		return
	}

	cleaned := sanitizeChirp(params.Body)

	parsedChirp := database.CreateChirpParams{
		Body:   cleaned,
		UserID: UserId,
	}
	respChirp, err := a.DB.CreateChirp(context.Background(), parsedChirp)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error storeing chirp in db: %s", err)
		log.Println(params.UserId)
		log.Println(UserId)
		return
	}

	crtdChirp := responseChirp{
		ID:        respChirp.ID,
		CreatedAt: respChirp.CreatedAt,
		UpdatedAt: respChirp.UpdatedAt,
		Body:      respChirp.Body,
		UserId:    respChirp.UserID,
	}

	respondWithJson(w, 201, crtdChirp)
}

func sanitizeChirp(chirp string) string {

	badWords := map[string]bool{"kerfuffle": true, "sharbert": true, "fornax": true}
	chWords := strings.Split(chirp, " ")

	for i, word := range chWords {
		if ok := badWords[strings.ToLower(word)]; ok {
			chWords[i] = "****"
		}
	}

	sanitized := strings.Join(chWords, " ")

	return sanitized
}
