package config

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (a *ApiConfig) UserHandler(w http.ResponseWriter, r *http.Request) {

	params := userReq{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil || params.Password == "" {
		respondWithError(w, 400, "bad request")
		return
	}

	hashed, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error hashing password: %s", err)
		return
	}

	dbUser := database.CreateUserParams{
		Email:    params.Email,
		Password: hashed,
	}

	newUser, err := a.DB.CreateUser(context.Background(), dbUser)
	if err != nil {
		respondWithError(w, 500, genericError)
		log.Printf("Error with user creation in DB: %s", err)
		return
	}

	parsedUser := userResp{
		Id:        newUser.ID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.CreatedAt,
		Email:     newUser.Email,
		IsRed:     newUser.IsChirpyRed,
	}

	respondWithJson(w, 201, parsedUser)
}
