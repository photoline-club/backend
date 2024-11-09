package auth

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		panic("Unable to hash") // TODO: is this recoverable?
	}
	return string(hash)
}

func PasswordValid(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func GenerateUID(length int) string {
	output := ""
	for i := 0; i < length; i++ {
		output += string(rune([]int{65, 97}[rand.Intn(2)] + rand.Intn(26)))
	}
	return output
}

func GenerateToken() string {
	return GenerateUID(32)
}
