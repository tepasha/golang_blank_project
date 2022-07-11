package repository

import (
	"fmt"

	pmcBackend "github.com/tepasha/golang_blank_project/internal/models"
)

type AuthPostgres struct {
	db *PostgresClient
}

func NewAuthPostgres(db *PostgresClient) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(user pmcBackend.User) (int64, error) {
	var userNumber int64

	query := fmt.Sprintf("INSERT INTO %s (email, phone, license, ulogin, pass, avatar, birthday, dtadd, dtupdate, about) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW(), $8, $9, $10) RETURNING id", tUser)
	row := p.db.Connection.QueryRow(query, user.Email, user.Phone, user.License, user.Login, user.PassWord, user.Avatar, user.Birthday, user.About)
	if err := row.Scan(&userNumber); err != nil {
		return 0, err
	}

	return userNumber, nil
}

func (p *AuthPostgres) GetUser(username, password string) (pmcBackend.User, error) {
	var user pmcBackend.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE ulogin = $1 AND pass = $2", tUser)
	err := p.db.Connection.Get(&user, query, username, password)

	return user, err
}
