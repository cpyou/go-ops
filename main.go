package main

import (
	"fmt"
	"go-ops/pkg/app"
	"go-ops/pkg/settings"
	"go-ops/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter() // router: gin.Engine
	app.AddAdmin(router)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}
