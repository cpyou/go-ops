package auth

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-ops/config"
	"time"
)

var (
	host, password string
)

func init() {
	c := config.GetConfig()
	host = c.Cache.Host
	password = c.Cache.Password
}

func RegisterSession() gin.HandlerFunc {
	store, _ := sessions.NewRedisStore(
		10,
		"tcp",
		host,
		password,
		[]byte("secret"))
	return sessions.Sessions("ops-session", store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(host,
		password, time.Minute)
	return cache.Cache(&cacheStore)
}
