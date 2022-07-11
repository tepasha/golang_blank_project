package service

import (
	"github.com/tepasha/golang_blank_project/internal/models"
	"github.com/tepasha/golang_blank_project/internal/repository"
)

type (
	Authorization interface {
		CreateUser(user models.User) (int64, error)
		GenerateToken(userLogin, userPassword string) (string, error)
		ParseToken(token string) (int64, error)
	}

	User interface {
		UpdateUser(userId int64, input models.UpdateUserInput) error
	}

	Service struct {
		Authorization
		User
	}
)

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
	}
}
