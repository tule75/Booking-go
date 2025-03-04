package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) (string, error) {
	var hash = sha256.New()

	hash.Write([]byte(key))

	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes), nil
}

// func HashPassword(password string, salt string) (string, error) {
// 	saltPassword := password + salt

// }
