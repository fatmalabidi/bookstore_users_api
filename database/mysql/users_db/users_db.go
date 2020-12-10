package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Client *sql.DB

const (
	sqlUsername = "sql_username"
	sqlPassword = "sql_password"
	sqlHost     = "sql_host"
	sqlPort     = "sql_port"
	schema      = "users_db"
)

func init() {
	username := os.Getenv(sqlUsername)
	password := os.Getenv(sqlPassword)
	port := os.Getenv(sqlPort)
	host := os.Getenv(sqlHost)
	// todo move to config and .env
	// dataSourceName= <userName>:<password>@tcp(<host>)/<schema>
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, host, port, schema)
	log.Println("connecting to ", dataSourceName)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
