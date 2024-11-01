package passwordhashing

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string, salt string) (string, error) {
	saltedPassword := password + salt
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func CheckPassword(password string, salt string, hashedPass string) bool {
	saltedPassword := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(saltedPassword))
	if err != nil {
		return false
	}
	return true
}
