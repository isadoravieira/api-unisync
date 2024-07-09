package security

import "golang.org/x/crypto/bcrypt"

// Hash takes a password and hashes it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// checkPassword compares a password and a hash and checks if they are the same
func checkPassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}