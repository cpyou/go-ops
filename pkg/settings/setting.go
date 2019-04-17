package settings

import (
	"go-ops/config"
	"log"
	"time"
)

var (
	RunMode  bool
	HTTPPort int

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret []byte
)

var c = config.GetConfig().Common

func init() {
	var err error
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}
	LoadServer()
	LoadApp()
}

func LoadServer() {
	RunMode = c.DEBUG
	HTTPPort = c.HTTP_PORT
	ReadTimeout = time.Duration(c.READ_TIMEOUT) * time.Second
	WriteTimeout = time.Duration(c.WRITE_TIMEOUT) * time.Second
}

func LoadApp() {
	secret := c.APP_SECRET
	JwtSecret = []byte(secret)
	PageSize = c.PAGE_SIZE
}
