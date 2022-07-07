package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Encrypt password
func Encrypt(pass string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return
}

// Decrypt password
func Decrypt(hashpass []byte, pass string) error {
	err := bcrypt.CompareHashAndPassword(hashpass, []byte(pass))
	if err != nil {
		return err
	}

	return err
}
