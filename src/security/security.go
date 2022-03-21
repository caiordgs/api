package security

import "golang.org/x/crypto/bcrypt"

// Hash receives a password string and hashes it.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// PasswordValidation validates the hashed password and the string password.
func PasswordValidation(hashPassword, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
}
