package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// User is an abstraction of a real-life user
type User struct {
	Username string
	Password string
}

// HashAndSalt hashes a raw string password to store it in the database safely
func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePlain := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
