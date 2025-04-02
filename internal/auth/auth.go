package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	pass := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(pass, 10)
	if err != nil {
		return "", err
	}
	passStr := string(hashed)
	return passStr, nil
}
