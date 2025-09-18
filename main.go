package main

import (
	"chirpy/internal/config"
	"chirpy/internal/database"
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	jwtSecret := os.Getenv("SECRET")
	polkaKey := os.Getenv("POLKA_KEY")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Error connecting to db: %s", err)
		os.Exit(1)
	}

	nMux := http.NewServeMux()
	apiCfg := config.ApiConfig{
		FileserverHits: atomic.Int32{},
		DB:             database.New(db),
		Secret:         jwtSecret,
		PolkaKey:       polkaKey,
	}
	appHandler := apiCfg.MetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	nMux.Handle("GET /app/", appHandler)
	nMux.HandleFunc("GET /admin/healthz", config.HelthHandler)
	nMux.HandleFunc("POST /admin/reset", apiCfg.ResetHandler)
	nMux.HandleFunc("POST /api/users", apiCfg.UserHandler)
	nMux.HandleFunc("POST /api/chirps", apiCfg.ChirpHandler)
	nMux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
	nMux.HandleFunc("GET /api/chirps", apiCfg.GetChirpsHandler)
	nMux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.GetAChirpHandler)
	nMux.HandleFunc("POST /api/login", apiCfg.LoginHandler)
	nMux.HandleFunc("POST /api/refresh", apiCfg.RefreshTokenHandler)
	nMux.HandleFunc("POST /api/revoke", apiCfg.RevokeHandler)
	nMux.HandleFunc("PUT /api/users", apiCfg.EmailPasswordHandler)
	nMux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.DeleteChirpHandler)
	nMux.HandleFunc("POST /api/polka/webhooks", apiCfg.UpdateRedHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: nMux,
	}

	server.ListenAndServe()
}
