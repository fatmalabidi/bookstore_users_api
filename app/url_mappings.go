package app

import (
	"github.com/fatmalabidi/bookstore_users_api/controllers/ping"
	"github.com/fatmalabidi/bookstore_users_api/controllers/users"
)

func mapUrls() {
 	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	// [host:port]/users&userID=<ID>
	router.GET("/users:userID", users.GetUser)

}
