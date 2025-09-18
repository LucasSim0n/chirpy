package yourpackage_test

import (
	"chirpy/internal/auth"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	secret := "supersecret"
	userID := uuid.New()
	expiresIn := time.Hour

	// Crear token
	token, err := auth.MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT devolvió error inesperado: %v", err)
	}
	if token == "" {
		t.Fatalf("MakeJWT devolvió un token vacío")
	}

	// Validar token
	parsedID, err := auth.ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT devolvió error inesperado: %v", err)
	}
	if parsedID != userID {
		t.Errorf("Se esperaba UserID %v, pero se obtuvo %v", userID, parsedID)
	}
}

func TestValidateJWT_InvalidSecret(t *testing.T) {
	secret := "correct_secret"
	wrongSecret := "wrong_secret"
	userID := uuid.New()
	expiresIn := time.Hour

	// Crear token válido
	token, err := auth.MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT devolvió error inesperado: %v", err)
	}

	// Validar con secret incorrecto
	_, err = auth.ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Errorf("Se esperaba error al validar con secret incorrecto, pero no ocurrió")
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	secret := "supersecret"
	userID := uuid.New()

	// Crear token expirado
	token, err := auth.MakeJWT(userID, secret, -time.Minute)
	if err != nil {
		t.Fatalf("MakeJWT devolvió error inesperado: %v", err)
	}

	// Validar token expirado
	_, err = auth.ValidateJWT(token, secret)
	if err == nil {
		t.Errorf("Se esperaba error por token expirado, pero no ocurrió")
	}
}

func TestValidateJWT_InvalidTokenString(t *testing.T) {
	secret := "supersecret"

	// Token inválido
	_, err := auth.ValidateJWT("not.a.jwt", secret)
	if err == nil {
		t.Errorf("Se esperaba error por token inválido, pero no ocurrió")
	}
}
