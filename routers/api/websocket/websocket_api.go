package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// websocket接口
func ViewEcsMonitor(c *gin.Context) {
	ws, err := Upgrade(c)
	if err != nil {
		fmt.Println(err)
	}
	go Writer(ws)
}

