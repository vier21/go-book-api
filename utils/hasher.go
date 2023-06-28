package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		return "", err
	}

	return string(byte), nil
}

func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}