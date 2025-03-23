package utils

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(hash), err
}

func ComparePass(pass string, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}
