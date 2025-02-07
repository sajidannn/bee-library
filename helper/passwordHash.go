package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 14
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", errors.New("hashing password failed")
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
