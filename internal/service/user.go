package service

import (
	"github.com/tepasha/golang_blank_project/internal/models"
	"github.com/tepasha/golang_blank_project/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpdateUser(userId int64, input models.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateUser(userId, input)
}
