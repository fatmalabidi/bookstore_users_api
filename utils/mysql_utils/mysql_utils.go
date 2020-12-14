package mysql_utils

import (
	errH "github.com/fatmalabidi/bookstore_users_api/utils/error_handler"
	"github.com/go-sql-driver/mysql"
	"strings"
)

func ParseError(err error) *errH.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errH.NewNotFoundError("user with the given ID not found")
		}
		return errH.NewInternalServerError("error parsing database error")
	}

	switch sqlErr.Number {
	case 1062:
		return errH.NewBadRequestError("duplicated data")
	default:
		return errH.NewInternalServerError("error processing request")

	}
}
