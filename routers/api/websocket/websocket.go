package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-ops/pkg/util"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

// 将http请求升级为websocket
func Upgrade(c *gin.Context) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, err
}

func Writer(conn *websocket.Conn) {

	for {
		ticker := time.NewTicker(5 * time.Second)

		for t := range ticker.C {
			fmt.Printf("Updating Stats: %+v\n", t)
			items, err := util.GetAsset()
			if err != nil {
				fmt.Println(err)
			}
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}