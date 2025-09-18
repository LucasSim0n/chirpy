package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost := 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
