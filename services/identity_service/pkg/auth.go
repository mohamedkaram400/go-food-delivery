package pkg

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func HashToken(token string) string {
    // Convert string to bytes and pass to Sum256
    hashBytes := sha256.Sum256([]byte(token))
	
	// Convert the [32]byte array into a readable hexadecimal string
	return hex.EncodeToString(hashBytes[:])
}

func CheckPassword(password string, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword(
        []byte(hashedPassword),
        []byte(password),
    )
}

