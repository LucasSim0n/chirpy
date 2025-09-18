package config

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (a *ApiConfig) EmailPasswordHandler(w http.ResponseWriter, r *http.Request) {
	t, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error getting jwt: %s", err)
		return
	}

	id, err := auth.ValidateJWT(t, a.Secret)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error validating jwt: %s", err)
		return
	}

	var body EmailPasswdUpdate

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error decoding json: %s", err)
		return
	}

	hash, err := auth.HashPassword(body.Password)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error hashing password: %s", err)
		return
	}

	dbUpdate := database.UpdateEmailPasswordParams{
		ID:        id,
		Email:     body.Email,
		Password:  hash,
		UpdatedAt: time.Now(),
	}

	updated, err := a.DB.UpdateEmailPassword(context.Background(), dbUpdate)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error updating db: %s", err)
		return
	}

	rUser := userResp{
		Id:           updated.ID,
		Email:        updated.Email,
		CreatedAt:    updated.CreatedAt,
		UpdatedAt:    updated.UpdatedAt,
		RefreshToken: updated.Token,
		Token:        t,
		IsRed:        updated.IsChirpyRed,
	}

	respondWithJson(w, 200, rUser)

}

type EmailPasswdUpdate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
