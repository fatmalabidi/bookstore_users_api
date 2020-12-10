package users

// the access layer  to db

import (
	"context"
	"fmt"
	"github.com/fatmalabidi/bookstore_users_api/database/mysql/users_db"
	"github.com/fatmalabidi/bookstore_users_api/utils/date_utils"
	errH "github.com/fatmalabidi/bookstore_users_api/utils/error_handler"
)

const (
	//  todo create builder/constructor to generate queries
	insertQuery = "INSERT INTO users(id,first_name, last_name, email, created_at, updated_at, date_created)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?);"
)

var (
	userDb = make(map[int64]*User)
)

func (user *User) Save() *errH.RestErr {
	statement, err := users_db.Client.PrepareContext(context.Background(), insertQuery)
	if err != nil {
		return errH.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = statement.Close()
	}()
	user.DateCreated = date_utils.GetNowString()
	user.CreatedAt = date_utils.GetNowSEpoch()
	user.UpdatedAt = date_utils.GetNowSEpoch()
	res, err := statement.Exec(user.ID, user.FirstName, user.LastName, user.Email, user.CreatedAt, user.UpdatedAt, user.DateCreated)
	if err != nil {
		return errH.NewInternalServerError(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		return errH.NewInternalServerError(err.Error())
	}
	user.ID = id
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
