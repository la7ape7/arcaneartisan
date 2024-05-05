package utils

import "golang.org/x/crypto/bcrypt"

func HashString(toHash string) (string, error) {
	hashed, hashError := bcrypt.GenerateFromPassword([]byte(toHash), bcrypt.DefaultCost)
	if hashError != nil {
		return "", hashError
	}
	return string(hashed), nil
}
