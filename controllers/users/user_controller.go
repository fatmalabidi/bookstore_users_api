package users

import (
	"github.com/fatmalabidi/bookstore_users_api/domain/users"
	"github.com/fatmalabidi/bookstore_users_api/services"
	resterr "github.com/fatmalabidi/bookstore_users_api/utils/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	//bytes, err := ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	// TODO handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	// TODO handle error
	//	return
	//}

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
	ctx.JSON(http.StatusNotImplemented, "GetUser: implement me")
}
