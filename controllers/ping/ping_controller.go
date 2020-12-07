package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// Ping is a function that will be used to test the connectivity to our server
func Ping(ctx *gin.Context){
	ctx.String(http.StatusOK,"pong")
}
