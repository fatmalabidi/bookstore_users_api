package users

import (
	errH "github.com/fatmalabidi/bookstore_users_api/utils/error_handler"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Password    string `json:"-"` //internal field, we don't use [Password] with json
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errH.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errH.NewBadRequestError("invalid password")

	}
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errH.NewBadRequestError("invalid email address")
	}
	return nil
}
