package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashpassword), err
}

func VerifyPassword(hashpassword string ,  candidatePassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(candidatePassword))

}