package checker

import (
	//"fmt"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
    bytes := sha256.Sum256([]byte(password))
	result := fmt.Sprint(bytes)
    return result
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}