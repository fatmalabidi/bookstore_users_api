package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "CreateUser: implement me")
}

func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "GetUser: implement me")
}
