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

	userID := getUserID(ctx)

	res, getErr := services.GetUser(userID)
	if getErr != nil {
		ctx.JSON(getErr.Code, getErr)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func UpdateUser(ctx *gin.Context) {
	userID := getUserID(ctx)

	var user users.User
	isPatch := ctx.Request.Method == http.MethodPatch

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := resterr.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Code, restErr)
		return
	}
	user.ID = userID

	res, updateErr := services.UpdateUser(user, isPatch)
	if updateErr != nil {
		ctx.JSON(updateErr.Code, updateErr)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func DeleteUser(ctx *gin.Context) {
	userID := getUserID(ctx)

	deleteErr := services.DeleteUser(userID)
	if deleteErr != nil {
		ctx.JSON(deleteErr.Code, deleteErr)
		return
	}
	ctx.JSON(http.StatusOK, "user has been successfully deleted")
}

func Search(ctx *gin.Context) {
	status := ctx.Query("status")
	userByStatus, err := services.Search(status)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}
	ctx.JSON(http.StatusOK, userByStatus)
}

func getUserID(ctx *gin.Context) int64 {
	userID, err := strconv.ParseInt(ctx.Param("userID"), 10, 64)
	if err != nil {
		restErr := resterr.NewBadRequestError("invalid param")
		ctx.JSON(restErr.Code, restErr)
		return 0
	}
	return userID
}
