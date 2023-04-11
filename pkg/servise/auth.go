package servise

import (
	"crypto/sha1"
	"fmt"
	"todo"
	"todo/pkg/repository"
)

const salt = "jdskds323ej"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generationPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generationPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
