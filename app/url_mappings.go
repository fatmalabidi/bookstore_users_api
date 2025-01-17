package app

import (
	"github.com/fatmalabidi/bookstore_users_api/controllers/ping"
	"github.com/fatmalabidi/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users:userID", users.GetUser)
	router.PUT("/users:userID", users.UpdateUser)
	router.PATCH("/users:userID", users.UpdateUser)
	router.DELETE("/users:userID", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)

}
