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

func (h *hasher) HashPassword(p string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

func (h *hasher) ComparePassword(p, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
