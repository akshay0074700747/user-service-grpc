package helpers

import "golang.org/x/crypto/bcrypt"

func Hash_pass(pass string) (string, error) {

	password := []byte(pass)

	hashedpass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(hashedpass), err

}

func VerifyPassword(password, checkpassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(password), []byte(checkpassword))

}
