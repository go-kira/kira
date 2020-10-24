package helpers

import "golang.org/x/crypto/bcrypt"

// BcryptHash - generate hashed password.
func BcryptHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// BcryptCompare - compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error on failure.
func BcryptCompare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
