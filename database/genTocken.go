package database

import (
	"crypto/rand"
	"encoding/hex"

)

type Token struct {
	Token string
	Registered string
}

func GenerateToken() string{
	b := make([]byte, 10)
    if _, err := rand.Read(b); err != nil {  
    }
	token := hex.EncodeToString(b)
	return token
}