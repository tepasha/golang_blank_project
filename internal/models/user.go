package models

import (
	"errors"
	"net/mail"
	"time"
)

type (
	FileStruct struct {
		Mime string `json:"mime"`
		Data string `json:"data"`
	}

	User struct {
		Id        int64      `json:"-" db:"id"`
		Email     string     `json:"email" db:"email" binding:"required"`
		Login     string     `json:"login" db:"ulogin" binding:"required"`
		Phone     string     `json:"phone" db:"phone"`
		License   string     `json:"license" db:"license"`
		PassWord  string     `json:"pass" db:"pass" binding:"required"`
		About     string     `json:"about" db:"about"`
		AvatarSrc FileStruct `json:"avatar" db:"-"`
		Avatar    string     `json:"-" db:"avatar"`
		Birthday  string     `json:"birthday" db:"birthday"`
	}

	UpdateUserInput struct {
		Email     string     `json:"email"`
		Login     string     `json:"login"`
		PassWord  string     `json:"pass"`
		About     string     `json:"about"`
		AvatarSrc FileStruct `json:"avatar" db:"-"`
		Avatar    string     `json:"-"`
		Birthday  string     `json:"birthday"`
	}
)

func (us *User) Validate() error {
	if us.Login == "" && len(us.Login) > 50 {
		return errors.New("invalid login")
	}

	if _, err := mail.ParseAddress(us.Email); err != nil {
		if len(us.Email) > 120 {
			return errors.New("invalid email")
		}
		return errors.New("invalid email")
	}

	if us.Birthday != "" {
		if birthday, err := time.Parse("2006-01-02", us.Birthday); err == nil {
			if birthday.After(time.Now()) {
				return errors.New("invalid birthday")
			}
		} else {
			return errors.New("invalid birthday")
		}
	} else {
		us.Birthday = "1800-02-01"
	}

	return nil
}

func (us *UpdateUserInput) Validate() error {
	if us.Login == "" && len(us.Login) > 50 {
		return errors.New("invalid login")
	}

	if _, err := mail.ParseAddress(us.Email); err != nil {
		if len(us.Email) > 120 {
			return errors.New("invalid email")
		}
		return errors.New("invalid email")
	}

	if birthday, err := time.Parse("2006-01-02", us.Birthday); err == nil {
		if birthday.After(time.Now()) {
			return errors.New("invalid birthday")
		}
	} else {
		return errors.New("invalid birthday")
	}

	return nil
}
