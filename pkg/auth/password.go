package auth

import (
	"barbershop/creativo/pkg/errors"
	"barbershop/creativo/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

// CompareHashAndPassword compares a hashed password with a plain text password.
func CompareHashAndPassword(hashedPassword types.HashedPassword, plainPassword types.Password) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return errors.ErrWrongCredentials
	}
	return nil
}

// HashPassword hashes a plain text password.
func HashPassword(plainPassword types.Password) (types.HashedPassword, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return types.HashedPassword(string(hash)), nil
}
