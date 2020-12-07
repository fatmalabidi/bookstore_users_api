package services

import (
	"github.com/fatmalabidi/bookstore_users_api/domain/users"
	resterr "github.com/fatmalabidi/bookstore_users_api/utils/error"
)

func CreateUser(user users.User) (*users.User, *resterr.RestErr) {
	return &user, nil
}
