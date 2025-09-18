package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	size := 32
	token := make([]byte, size)
	_, err := rand.Read(token)
	return hex.EncodeToString(token), err
}
