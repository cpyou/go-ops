package auth_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
	"go-ops/models/users"
	"go-ops/pkg/app"
	"go-ops/pkg/apierror"
	"go-ops/pkg/util"
)

func ViewGetToken(c *gin.Context) {
	appG := app.Gin{C: c}
	var user users.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		appG.FailResponse(apierror.ERROR_AUTH)
	} else {
		if r := util.CheckUser(user.Username, user.Password); r {
			token, _ := util.GenerateToken(user.Username, user.Password)
			appG.SuccessResponse(gin.H{"token": token})
		} else {
			appG.FailResponse(apierror.ERROR_AUTH)
		}
	}
}

func ViewCreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var user users.User
	var err error
	u, err := appG.GetBodyData(user) // u map[string]interface{}
	mapstructure.Decode(u, &user)
	if err != nil {
		fmt.Println(err)
		appG.FailResponse(apierror.INVALID_PARAMS)
	} else {
		if err := util.CreateUser(user); err != nil {
			appG.FailResponse(apierror.INVALID_PARAMS)
		} else {
			appG.SuccessResponse("创建成功")
		}
	}
}

func ChangePassword(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}

func ChangeProfile(c *gin.Context) {

}
