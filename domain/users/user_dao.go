package users

// the access layer  to db

import (
	"fmt"
	errH "github.com/fatmalabidi/bookstore_users_api/utils/error_handler"
)

var (
	userDb = make(map[int64]*User)
)

func (user *User) Save() *errH.RestErr {
	currentUser := userDb[user.ID]
	if currentUser != nil {
		return errH.NewBadRequestError(fmt.Sprintf("user with id %d already exist", user.ID))
	}

	userDb[user.ID] = user
	return nil
}

func (user *User) Get() *errH.RestErr {
	currentUser := userDb[user.ID]
	if currentUser == nil {
		return errH.NewNotFoundError(fmt.Sprintf("user with id %d not found", user.ID))
	}
	user.ID = currentUser.ID
	user.CreatedAt = currentUser.CreatedAt
	user.UpdatedAt = currentUser.UpdatedAt
	user.Email = currentUser.Email
	user.DateCreated = currentUser.DateCreated
	user.LastName = currentUser.LastName
	user.FirstName = currentUser.FirstName
	return nil
}
