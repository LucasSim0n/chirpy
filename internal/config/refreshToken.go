package config

import (
	"chirpy/internal/auth"
	"context"
	"log"
	"net/http"
	"time"
)

func (a *ApiConfig) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error getting refresh token: %s", err)
		return
	}

	dbToken, err := a.DB.GetRefreshToken(context.Background(), token)
	if err != nil || dbToken.ExpiresAt.Before(time.Now()) || dbToken.RevokedAt.Valid != false {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error getting refresh token from db: %s", err)
		return
	}

	authToken, err := auth.MakeJWT(dbToken.UserID, a.Secret)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error creationg JWT: %s", err)
		return
	}

	type jsonToken struct {
		Token string `json:"token"`
	}

	jtoken := jsonToken{authToken}

	respondWithJson(w, 200, jtoken)
}
