package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(UserID uuid.UUID, tokenSecret string) (string, error) {

	expiresIn := time.Duration(time.Hour)
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(expiresIn)),
		Subject:   UserID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte(tokenSecret))
	return str, err
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {

	claims := jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) { return []byte(tokenSecret), nil })
	if err != nil {
		return uuid.UUID{}, err
	}

	stringId, err := token.Claims.GetSubject()
	if err != nil {
		log.Println("Error getting subject")
		return uuid.UUID{}, err
	}
	id, err := uuid.Parse(stringId)
	if err != nil {
		log.Println("Error parsing string to uuid")
		return uuid.UUID{}, err
	}

	return id, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	fullAuth, ok := headers[http.CanonicalHeaderKey("Authorization")]
	if !ok {
		return "", errors.New("Authorization header not provided.")
	}

	auth := strings.Split(fullAuth[0], " ")
	if auth[0] != "Bearer" {
		return "", errors.New("Bearer token not provided")
	}

	return auth[1], nil
}
