package config

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
)

func (a *ApiConfig) RevokeHandler(w http.ResponseWriter, r *http.Request) {

	refresh, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		log.Printf("Error getting refresh token: %s", err)
		return
	}

	tokenToRevoke := database.RevokeTokenParams{
		Token: refresh,
		RevokedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true},
		UpdatedAt: time.Now(),
	}

	a.DB.RevokeToken(context.Background(), tokenToRevoke)
	w.WriteHeader(204)

}
