package hasher

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Manager interface {
	HashPassword(password string) string
	ComparePassword(password, hash string) bool
}

type hasher struct {
}

func CreateManager() Manager {
	return &hasher{}
}

func (h *hasher) HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

func (h *hasher) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
