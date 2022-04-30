package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher provides hashing logic to securely store passwords.
type PasswordHasher interface {
	Hash(password string) (string, error)
	IsEqual(pwdHashed, pwd string) bool
}

// SHA256Hasher uses SHA1 to hash passwords with provided salt.
type SHA256Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{salt: salt}
}

// Hash creates SHA1 hash of given password.
func (h *SHA256Hasher) Hash(password string) (string, error) {
	pwdByte := []byte(password)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(pwdByte, 11)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *SHA256Hasher) IsEqual(pwdHashed, pwd string) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(pwdHashed), []byte(pwd))
	if err != nil {
		fmt.Println(err) // nil means it is a match
		return false
	}
	return true
}
