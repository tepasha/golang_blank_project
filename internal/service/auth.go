package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	models "github.com/tepasha/golang_blank_project/internal/models"
	"github.com/tepasha/golang_blank_project/internal/repository"
)

const (
	salt      = "fuckthemall"
	tokenTTL  = 12 * time.Hour
	singInKey = "singinkeymother"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int64, error) {
	if err := user.Validate(); err != nil {
		return 0, err
	}
	user.PassWord = generatePasswordHash(user.PassWord)
	return s.repo.CreateUser(user)
}
