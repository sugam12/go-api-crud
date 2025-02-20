package auth

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(password1 string, password2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	if err != nil {
		return err
	}
	return nil
}
