package users

import (
	"github.com/fatmalabidi/bookstore_users_api/domain/users"
	"github.com/fatmalabidi/bookstore_users_api/services"
	resterr "github.com/fatmalabidi/bookstore_users_api/utils/error_handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := resterr.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Code, restErr)
		return
	}
	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Code, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func GetUser(ctx *gin.Context) {

	userID, err := strconv.ParseInt(ctx.Param("userID"), 10, 64)
	if err != nil {
		restErr := resterr.NewBadRequestError("invalid param")
		ctx.JSON(restErr.Code, restErr)
		return
	}

	res, getErr := services.GetUser(userID)
	if getErr != nil {
		ctx.JSON(getErr.Code, getErr)
		return
	}
	ctx.JSON(http.StatusOK, res)}
