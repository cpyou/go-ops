package app

import (
	_ "github.com/chenhg5/go-admin/adapter/gin" // 引入adapter
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/gin-gonic/gin"
	"go-ops/pkg/app/generators"
)

func AddAdmin(router *gin.Engine) {
	// 加载admin
	eng := engine.Default()
	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "zs",  // database
				PWD:          "zs",
				NAME:         "go_ops",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: "127.0.0.1:8009", // the domain of cookie which be used when visiting your site.
		PREFIX: "admin",
		// STORE is important. And the directory should has permission to write.
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE: "cn",
		INDEX:    "/",
	}
	adminPlugin := admin.NewAdmin(generators.Generators)
	examplePlugin := example.NewExample()

	eng.AddConfig(cfg).
		AddPlugins(adminPlugin, examplePlugin).  // 加载插件
		Use(router)
}
