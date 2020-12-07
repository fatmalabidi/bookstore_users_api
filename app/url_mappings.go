package app

import (
	"github.com/fatmalabidi/bookstore_users_api/controllers/ping"
	"github.com/fatmalabidi/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users:userID", users.CreateUser)
	router.GET("/users", users.GetUser)

}
