package routers

import (
	_ "github.com/chenhg5/go-admin/adapter/gin" // 必须引入，如若不引入，则需要自己定义
	"github.com/gin-gonic/gin"
	"go-ops/middlewares/auth"
	"go-ops/middlewares/erroremail"
	"go-ops/middlewares/permission"
	"go-ops/routers/api/auth_api"
	"go-ops/routers/api/v1"
	"go-ops/routers/api/websocket"
)

func InitRouter() *gin.Engine {
	// 日志与恢复中间件  gin.New() 不带这两个中间件
	r := gin.Default()
	r.Use(erroremail.ErrEmailWriter())  // 500发送报错邮件

	apiAuth := r.Group("api/auth")
	{
		apiAuth.POST("/create-user/", auth.AuthMiddleware([]string{"cookie", "jwt"}),
			permission.PermissionMiddleware(),
			auth_api.ViewCreateUser)
		apiAuth.GET("/get-token/", auth_api.ViewGetToken)
	}

	apiv1 := r.Group("api/v1")
	apiv1.Use(auth.AuthMiddleware([]string{"cookie", "jwt"})) // 认证中间件, 默认先cookie后jwt
	apiv1.Use(permission.PermissionMiddleware())              // 鉴权中间件

	{
		apiv1.GET("/asset/:id/", v1.ViewGetAsset)
	}

	apiws := r.Group("ws/v1")
	{
		apiws.GET("/ecs-monitor/", websocket.ViewEcsMonitor)
	}

	return r
}