package users

// the access layer  to db

import (
	"fmt"
	"github.com/fatmalabidi/bookstore_users_api/database/mysql/users_db"
	dateUtils "github.com/fatmalabidi/bookstore_users_api/utils/date_utils"
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
	user.CreatedAt = dateUtils.GetNowSEpoch()
	user.UpdatedAt = dateUtils.GetNowSEpoch()
	user.DateCreated = dateUtils.GetNowString()
	userDb[user.ID] = user
	return nil
}

func (user *User) Get() *errH.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
