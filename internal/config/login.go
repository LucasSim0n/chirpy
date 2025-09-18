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

func (a *ApiConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := userReq{}

	err := decoder.Decode(&user)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error decoding request: %s", err)
		return
	}

	dbUser, err := a.DB.GetUserByEmail(context.Background(), user.Email)
	if err != nil {
		respondWithError(w, 401, "Incorrect email or password")
		log.Printf("Error getting user from db: %s", err)
		return
	}

	if err := auth.CheckPasswordHash(user.Password, dbUser.Password); err != nil {
		respondWithError(w, 401, "Incorrect email or password")
		log.Printf("Error checking password: %v", err)
		return
	}

	jtoken, err := auth.MakeJWT(dbUser.ID, a.Secret)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error generating jwt: %s", err)
		return
	}

	rtoken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error generating refresh token: %s", err)
		return
	}

	dbtoken := database.CreateRefreshTokenParams{
		Token:     rtoken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    dbUser.ID,
		ExpiresAt: time.Now().Add(time.Duration(60 * 24 * time.Hour)),
	}
	_, err = a.DB.CreateRefreshToken(context.Background(), dbtoken)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error creating refresh token: %s", err)
		return
	}

	userInfo := userResp{
		Id:           dbUser.ID,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		Email:        dbUser.Email,
		Token:        jtoken,
		RefreshToken: rtoken,
		IsRed:        dbUser.IsChirpyRed,
	}
	respondWithJson(w, 200, userInfo)
}
