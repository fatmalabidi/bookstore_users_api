package services

import (
	"fmt"
	"github.com/fatmalabidi/bookstore_users_api/domain/users"
	errH "github.com/fatmalabidi/bookstore_users_api/utils/error_handler"
)

func GetUser(userID int64) (*users.User, *errH.RestErr) {
	if userID < 0 {
		return nil, errH.NewBadRequestError("invalid userID")
	}
	res := &users.User{ID: userID}
	err := res.Get()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateUser(user users.User) (*users.User, *errH.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errH.RestErr) {
	currentUser := user
	if err := currentUser.Get(); err != nil {
		return nil, errH.NewNotFoundError(fmt.Sprintf("user with the id %d not found", user.ID))
	}
	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email
	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return &currentUser, nil
}
