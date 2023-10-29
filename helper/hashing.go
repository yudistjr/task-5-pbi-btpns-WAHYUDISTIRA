package helper

import "golang.org/x/crypto/bcrypt"

func PasswordHash(pw string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(result), err
}

func PasswordCheck(pw, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}