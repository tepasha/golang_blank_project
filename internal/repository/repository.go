package repository

import (
	"github.com/tepasha/golang_blank_project/internal/models"
)

type (
	Authorization interface {
		CreateUser(user models.User) (int64, error)
		GetUser(username, password string) (models.User, error)
	}

	User interface {
		UpdateUser(userId int64, input models.UpdateUserInput) error
	}

	Repository struct {
		Authorization
		User
	}
)

func NewRepository(db *PostgresClient) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
	}
}
