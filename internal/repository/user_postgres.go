package repository

import (
	"fmt"
	"strings"

	"github.com/tepasha/golang_blank_project/internal/models"
)

type UsersPostgres struct {
	db *PostgresClient
}

func NewUserPostgres(db *PostgresClient) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (up *UsersPostgres) UpdateUser(userId int64, input models.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if input.Email != "" {
		setValues = append(setValues, fmt.Sprintf("email = $%d", argsId))
		args = append(args, input.Email)
		argsId++
	}

	if input.PassWord != "" {
		setValues = append(setValues, fmt.Sprintf("pass = $%d", argsId))
		args = append(args, input.PassWord)
		argsId++
	}

	if input.Login != "" {
		setValues = append(setValues, fmt.Sprintf("ulogin = $%d", argsId))
		args = append(args, input.Login)
		argsId++
	}

	if input.About != "" {
		setValues = append(setValues, fmt.Sprintf("about = $%d", argsId))
		args = append(args, input.About)
		argsId++
	}

	if input.Avatar != "" {
		setValues = append(setValues, fmt.Sprintf("avatar = $%d", argsId))
		args = append(args, input.Avatar)
		argsId++
	}

	if input.Birthday != "" {
		setValues = append(setValues, fmt.Sprintf("birthday = $%d", argsId))
		args = append(args, input.Birthday)
		argsId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET dtupdate = now(), %s"+
		"WHERE id = $%d",
		tUser, setQuery, userId)

	args = append(args, userId)

	_, err := up.db.Connection.Exec(query, args...)

	return err
}
