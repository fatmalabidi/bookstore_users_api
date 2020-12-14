package users

// the access layer  to db

import (
	"context"
	"fmt"
	"github.com/fatmalabidi/bookstore_users_api/database/mysql/users_db"
	"github.com/fatmalabidi/bookstore_users_api/utils/date_utils"
	errH "github.com/fatmalabidi/bookstore_users_api/utils/error_handler"
	"github.com/fatmalabidi/bookstore_users_api/utils/mysql_utils"
)

const (
	//  todo create builder/constructor to generate queries
	insertQuery = "INSERT INTO users(id,first_name, last_name, email, password, status, created_at, updated_at, date_created)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);"

	getQuery = "SELECT  id,first_name, last_name, email, status, created_at, updated_at, date_created FROM users WHERE id=?"

	updateQuery = "UPDATE users SET first_name=? , last_name=?, email=?, updated_at=? WHERE id=?"

	deleteQuery = "DELETE FROM users WHERE id=?"

	findByStatusQuery = "SELECT  id,first_name, last_name, email, created_at, updated_at, date_created, status FROM users WHERE status=?"
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
	res, err := statement.Exec(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
		user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
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

	statement, err := users_db.Client.PrepareContext(context.Background(), getQuery)
	if err != nil {
		return errH.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = statement.Close()
	}()
	res := statement.QueryRow(user.ID)
	//  the scan dest should be pointers !
	if scanErr := res.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DateCreated,
	)
		scanErr != nil {
		return mysql_utils.ParseError(scanErr)
	}
	return nil
}

func (user *User) Update() *errH.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	statement, err := users_db.Client.PrepareContext(context.Background(), updateQuery)
	if err != nil {
		return errH.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = statement.Close()
	}()
	user.UpdatedAt = date_utils.GetNowSEpoch()
	_, err = statement.Exec(user.FirstName, user.LastName, &user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errH.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	statement, err := users_db.Client.PrepareContext(context.Background(), deleteQuery)
	if err != nil {
		return errH.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = statement.Close()
	}()
	user.UpdatedAt = date_utils.GetNowSEpoch()
	_, err = statement.Exec(user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) GetByStatus(status string) ([]User, *errH.RestErr) {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	statement, err := users_db.Client.PrepareContext(context.Background(), findByStatusQuery)
	if err != nil {
		return nil, errH.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = statement.Close()
	}()

	rows, err := statement.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer func() {
		_ = rows.Close()
	}()

	result := make([]User, 0)
	for rows.Next() {
		var u = User{}
		if err := rows.Scan(
			&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.DateCreated, &u.Status,
		); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, u)
	}
	if len(result) == 0 {
		return nil, errH.NewNotFoundError(fmt.Sprintf("no users matching the status %s", status))
	}
	return result, nil
}
